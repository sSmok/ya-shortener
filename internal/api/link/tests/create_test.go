package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/sSmok/ya-shortener/internal/api/link"
	"github.com/sSmok/ya-shortener/internal/repository"
	"github.com/sSmok/ya-shortener/internal/repository/mocks"
	"github.com/stretchr/testify/require"
)

func TestAPI_Create(t *testing.T) {
	type linkRepositoryMockFunc func(mc *minimock.Controller) repository.LinkRepository

	var (
		minimockContr = minimock.NewController(t)
	)

	tests := []struct {
		name               string
		method             string
		body               string
		expectedStatus     int
		expectedBody       string
		linkRepositoryMock linkRepositoryMockFunc
	}{
		{
			name:           "Valid POST request",
			method:         http.MethodPost,
			body:           "http://example.com",
			expectedStatus: http.StatusCreated,
			expectedBody:   "http://localhost:8080/shortURL",
			linkRepositoryMock: func(mc *minimock.Controller) repository.LinkRepository {
				mock := mocks.NewLinkRepositoryMock(mc)
				mock.CreateMock.Return("shortURL", nil)
				return mock
			},
		},
		{
			name:           "Invalid method",
			method:         http.MethodGet,
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "",
			linkRepositoryMock: func(mc *minimock.Controller) repository.LinkRepository {
				mock := mocks.NewLinkRepositoryMock(mc)
				return mock
			},
		},
		{
			name:           "Empty body",
			method:         http.MethodPost,
			body:           "",
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   "",
			linkRepositoryMock: func(mc *minimock.Controller) repository.LinkRepository {
				mock := mocks.NewLinkRepositoryMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := link.NewAPI(tt.linkRepositoryMock(minimockContr))

			req := httptest.NewRequest(tt.method, "/create", strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			api.Create(w, req)

			resp := w.Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Fatalf("failed to close response body: %v", err)
				}
			}(resp.Body)

			require.Equal(t, tt.expectedStatus, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed to read response body: %v", err)
			}

			require.Equal(t, tt.expectedBody, string(body))
		})
	}
}
