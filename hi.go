package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// HiHandler implements http.Handler
type HiHandler struct {
	config *Config
	logger *log.Logger
}

// ServeHTTP response /hi
func (h *HiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	client := RealIP(r)
	user := r.FormValue("name")
	h.logger.Printf("from: %s, name: %s", client, user)
	time.Sleep(time.Duration(h.config.SleepTime) * time.Second)

	ret := fmt.Sprintf("Hi %s From %s", user, client)
	http.Error(w, ret, http.StatusOK)
}
