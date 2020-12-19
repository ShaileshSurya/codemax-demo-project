package main

import (
	"net/http"
)

func xAPIAuthenticationMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	xAPIKeyHeader := r.Header.Get("x-api-key")
	if xAPIKey != xAPIKeyHeader {
		errMsg := Response{
			Message: "Invalid API Key",
		}
		writeResponse(r.Context(), w, errMsg, http.StatusUnauthorized)
		return
	}
	next(w, r)
}
