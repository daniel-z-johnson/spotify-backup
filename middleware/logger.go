package middleware

import (
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

func Logger(logger zerolog.Logger) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			logger.Info().
				Str("duration", time.Since(start).String()).
				Int("code", wrapped.Status()).
				Int("bytes", wrapped.TotalBytes).
				Str("path", r.URL.Path).
				Str("method", r.Method).
				Msg("request info")
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
