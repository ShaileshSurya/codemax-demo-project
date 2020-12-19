// request_logger.go
package main

import (
	"net/http"
	"runtime/debug"
	"time"

	//log "github.com/go-kit/kit/log"

	"github.com/sirupsen/logrus"
)

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

func loggerMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			log.WithFields(logrus.Fields{
				"err":   err,
				"trace": debug.Stack(),
			}).Info("Error!")
		}
	}()

	start := time.Now()
	wrapped := wrapResponseWriter(w)

	next(wrapped, r)

	log.WithFields(logrus.Fields{
		"status":   wrapped.status,
		"method":   r.Method,
		"path":     r.URL.EscapedPath(),
		"duration": time.Since(start),
	}).Info("Request Log dump")
}
