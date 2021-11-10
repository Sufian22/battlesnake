package middlewares

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			w.Header().Set("Content-Type", "application/json")
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)

			logger.WithFields(logrus.Fields{
				"status":  wrapped.status,
				"method":  r.Method,
				"path":    r.URL.EscapedPath(),
				"latency": time.Since(start),
			}).Info()
		}

		return http.HandlerFunc(fn)
	}
}
