package middleware

import (
	"net/http"
	"time"

	"github.com/sSmok/ya-shortener/internal/logger"
	"go.uber.org/zap"
)

type responseData struct {
	status int
	size   int
}

type loggingResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

func (w *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.responseData.size += size

	return size, err
}

func (w *loggingResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.responseData.status = statusCode
}

// Log middleware для логирования запросов
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		respData := &responseData{
			status: 0,
			size:   0,
		}
		lw := &loggingResponseWriter{w, respData}

		next.ServeHTTP(lw, r)

		logger.Info("request",
			zap.String("URI", r.RequestURI),
			zap.String("method", r.Method),
			zap.Int("status", respData.status),
			zap.Duration("duration", time.Since(now)),
			zap.Int("size", respData.size),
		)
	})
}
