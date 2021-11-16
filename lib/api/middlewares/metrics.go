package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sufian22/battlesnake/pkg/server/metrics"
)

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := wrapResponseWriter(w)
		next.ServeHTTP(wrapped, r)

		path := r.URL.EscapedPath()
		metrics.HttpRequestsDuration.WithLabelValues(fmt.Sprintf("%v", wrapped.status), r.Method, path).
			Observe(float64(time.Since(start)))
	})
}
