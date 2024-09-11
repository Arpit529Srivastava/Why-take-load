package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// Server interface defines methods for a load balancer to interact with backend servers
type Server interface {
	Address() string            // Method to get the server address
	IsAlive() bool              // Method to check if the server is alive
	Serve(rw http.ResponseWriter, r *http.Request) // Method to serve the HTTP request
}
// simpleserver struct represents a simple backend server with an address and reverse proxy
type simpleserver struct {
	addr  string                 // Server address
	proxy httputil.ReverseProxy   // Reverse proxy to forward the requests
}