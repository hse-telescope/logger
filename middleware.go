package logger

import "net/http"

func AddLoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			Info(r.Context(), "got request", "request", r)
			handler.ServeHTTP(w, r)
		},
	)
}
