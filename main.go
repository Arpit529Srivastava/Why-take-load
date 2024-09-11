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