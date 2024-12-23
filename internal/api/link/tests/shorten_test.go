package tests

import (
	"fmt"
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

func TestAPI_Shorten(t *testing.T) {
	type linkRepositoryMockFunc func(mc *minimock.Controller) repository.LinkRepository

	var (
		minimockContr = minimock.NewController(t)
		baseURL       = "http://localhost:8080"
		shortID       = "shortURL"
		respBody      = fmt.Sprintf(`{"result":"%s/%s"}`, baseURL, shortID)
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
			body:           `{"url": "https://practicum.yandex.ru"}`,
			expectedStatus: http.StatusCreated,
			expectedBody:   respBody,
			linkRepositoryMock: func(mc *minimock.Controller) repository.LinkRepository {
				mock := mocks.NewLinkRepositoryMock(mc)
				mock.CreateMock.Return(shortID, nil)
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
		{
			name:           "Broken JSON",
			method:         http.MethodPost,
			body:           `{"url": "https://practicum.yandex.ru}`,
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
			api := link.NewAPI(tt.linkRepositoryMock(minimockContr), baseURL)

			req := httptest.NewRequest(tt.method, "/create", strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			api.Shorten(w, req)

			resp := w.Result()
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Fatalf("failed to close response body: %v", err)
				}
			}(resp.Body)

			require.Equal(t, tt.expectedStatus, resp.StatusCode, "Response code didn't match expected")

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed to read response body: %v", err)
			}

			if tt.expectedBody != "" {
				require.JSONEq(t, tt.expectedBody, string(body))
			}
		})
	}
}
