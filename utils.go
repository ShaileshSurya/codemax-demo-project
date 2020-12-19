package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func writeResponse(ctx context.Context, w http.ResponseWriter, response Response, statusCode int) {
	byteArray, _ := json.Marshal(response)
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteArray)
}
