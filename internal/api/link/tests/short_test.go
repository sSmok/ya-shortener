package tests

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/sSmok/ya-shortener/internal/api/link"
	"github.com/sSmok/ya-shortener/internal/repository"
	"github.com/sSmok/ya-shortener/internal/repository/mocks"
	"github.com/stretchr/testify/require"
)

func TestAPI_Short(t *testing.T) {
	type linkRepositoryMockFunc func(mc *minimock.Controller) repository.LinkRepository

	var (
		minimockContr = minimock.NewController(t)
		originalURL   = "http://example.com"
		repoErr       = errors.New("repository error")
		baseURL       = "http://localhost:8080/"
	)

	tests := []struct {
		name               string
		shortID            string
		method             string
		expectedStatus     int
		expectedHeader     string
		linkRepositoryMock linkRepositoryMockFunc
	}{
		{
			name:           "Valid short ID",
			shortID:        "shortURL",
			method:         http.MethodGet,
			expectedStatus: http.StatusTemporaryRedirect,
			expectedHeader: originalURL,
			linkRepositoryMock: func(mc *minimock.Controller) repository.LinkRepository {
				mock := mocks.NewLinkRepositoryMock(mc)
				mock.GetMock.Return(originalURL, nil)
				return mock
			},
		},
		{
			name:           "Invalid method",
			shortID:        "invalid",
			method:         http.MethodGet,
			expectedStatus: http.StatusNotFound,
			expectedHeader: "",
			linkRepositoryMock: func(mc *minimock.Controller) repository.LinkRepository {
				mock := mocks.NewLinkRepositoryMock(mc)
				mock.GetMock.Return("", repoErr)
				return mock
			},
		},
		{
			name:           "Invalid short ID",
			shortID:        "shortURL",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedHeader: "",
			linkRepositoryMock: func(mc *minimock.Controller) repository.LinkRepository {
				mock := mocks.NewLinkRepositoryMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := link.NewAPI(tt.linkRepositoryMock(minimockContr), baseURL)

			req := httptest.NewRequest(tt.method, "/"+tt.shortID, nil)
			w := httptest.NewRecorder()

			api.Short(w, req)

			resp := w.Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Fatalf("failed to close response body: %v", err)
				}
			}(resp.Body)

			require.Equal(t, tt.expectedStatus, resp.StatusCode)

			location := resp.Header.Get("Location")
			require.Equal(t, tt.expectedHeader, location)
		})
	}
}
