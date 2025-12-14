package logger

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func New(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			entry := log.With(
				slog.String("request.method", r.Method),
				slog.String("request.host", r.Host),
				slog.String("request.url", r.RequestURI),
				slog.String("request.user_ip", r.RemoteAddr),
				slog.String("request.user_agent", r.UserAgent()),
			)

			entry.Debug(fmt.Sprintf("Routing: %s", r.RequestURI))

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			defer func() {
				entry.Debug(fmt.Sprintf("Server handled: %s", r.RequestURI))
			}()
			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
