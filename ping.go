package main

import "net/http"

// PingHandler implements http.Handler
type PingHandler struct{}

// ServeHTTP return server status
func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "pong", http.StatusOK)
}
