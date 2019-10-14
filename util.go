package main

import (
	"net"
	"net/http"
	"strings"
)

// RealIP get real client ip
// first from header: X-Forwarded-For, then from peer
func RealIP(r *http.Request) string {
	ipList := strings.SplitN(r.Header.Get("X-Forwarded-For"), ",", 2)
	ip := strings.TrimSpace(ipList[0])
	if ip != "" {
		return ip
	}

	ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	return ip
}
