package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

// HiHandler implements http.Handler
type HiHandler struct {
	logger *log.Logger
}

// ServeHTTP response /hi
func (h *HiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	user := r.FormValue("name")
	h.logger.Printf("from: %s, name: %s", host, user)

	w.Write([]byte(fmt.Sprintf("Hi %s From %s\n", user, host)))
}
