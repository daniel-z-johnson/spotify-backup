package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Logger(logger *slog.Logger) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			logger.Info("request info",
				slog.Duration("duration", time.Since(start)),
				slog.Any("code", wrapped.Status()),
				slog.Any("bytes", wrapped.TotalBytes),
				slog.Any("path", r.URL.Path),
				slog.Any("method", r.Method),
				slog.Any("requestAddr", r.RemoteAddr),
			)
		}

		return http.HandlerFunc(fn)
	}

}

type responseWriter struct {
	http.ResponseWriter
	code        int
	wroteHeader bool
	TotalBytes  int
}

func wrapResponseWriter(writer http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: writer}
}

func (rw *responseWriter) Status() int {
	if !rw.wroteHeader {
		rw.code = 200
	}
	return rw.code
}

func (rw *responseWriter) WriteHeader(code int) {
	if !rw.wroteHeader {
		rw.code = code
		rw.ResponseWriter.WriteHeader(code)
		rw.wroteHeader = true
	}
}

func (rw *responseWriter) Write(bytes []byte) (int, error) {
	rw.TotalBytes += len(bytes)
	return rw.ResponseWriter.Write(bytes)
}
