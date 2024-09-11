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

// Constructor function for creating a new simpleserver instance
func newsimplesever(addr string) *simpleserver {
	serverurl, err := url.Parse(addr) // Parse the server address as a URL
	handlErr(err)                     // Handle any errors that occur during parsing
	return &simpleserver{
		addr:  addr,                                // Set the address
		proxy: *httputil.NewSingleHostReverseProxy(serverurl), // Create a new reverse proxy to the given server
	}
}

// LoadBalancer struct contains the logic for distributing requests to servers
type LoadBalancer struct {
	port       string   // Port where the load balancer is running
	roundrobin int      // Index to keep track of which server to use next
	servers    []Server // Slice of servers that the load balancer will distribute traffic to
}
// NewLoadBalancer constructor initializes a LoadBalancer with a given port and list of servers
func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:       port,
		roundrobin: 0,      // Start with the first server in the list
		servers:    servers, // Set the list of backend servers
	}
}
func (s *simpleserver) Address() string { 
	return s.addr 
}
func (s *simpleserver) IsAlive() bool {
	return true
}
// Serve method forwards the HTTP request to the backend server using the reverse proxy
func (s *simpleserver) Serve(rw http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(rw, r) // Use the reverse proxy to handle the request
}
// getNextAvailableServer method implements round-robin selection of servers
func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundrobin%len(lb.servers)] // Select the next server in the list
	for !server.IsAlive() {  // Check if the server is alive; if not, move to the next one
		lb.roundrobin++
		server = lb.servers[lb.roundrobin%len(lb.servers)]
	}
	lb.roundrobin++ // Move to the next server for future requests
	return server   // Return the available server
}
// serveProxy method forwards the request to the selected backend server
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()  // Get the next available server
	fmt.Printf("forwarding request to address %q \n", targetServer.Address()) // Log the target server
	targetServer.Serve(rw, r) // Forward the request to the selected server
}

// handlErr function handles any errors, printing the error message and exiting the program
func handlErr(err error) {
	if err != nil {
		fmt.Println("error has occurred", err)
		os.Exit(1) // Exit if an error occurs
	}
}