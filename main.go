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
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

// simpleserver struct represents a simple backend server with an address and reverse proxy
type simpleserver struct {
	addr  string
	proxy httputil.ReverseProxy //reverse proxy is used for forwarding the requests to the server
}

// Constructor function for creating a new simpleserver instance
func newsimplesever(addr string) *simpleserver {
	serverurl, err := url.Parse(addr)
	handlErr(err)
	return &simpleserver{
		addr:  addr,
		proxy: *httputil.NewSingleHostReverseProxy(serverurl),
	}
}

// LoadBalancer struct contains the logic for distributing requests to servers
type LoadBalancer struct {
	port       string
	roundrobin int
	servers    []Server
}

// NewLoadBalancer constructor initializes a LoadBalancer with a given port and list of servers
func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:       port,
		roundrobin: 0,
		servers:    servers,
	}
}
func (s *simpleserver) Address() string {
	return s.addr
}
func (s *simpleserver) IsAlive() bool {
	resp, err := http.Get(s.addr)             // Send a simple GET request to the server
	if err != nil || resp.StatusCode >= 500 { // If there's an error or server error (5xx)
		return false // Server is not alive
	}
	return true // Server is alive
}

// Serve method forwards the HTTP request to the backend server using the reverse proxy
func (s *simpleserver) Serve(rw http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(rw, r)
}

// logic for selection of next availabe server
// MAIN LOGIC
func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundrobin%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundrobin++
		server = lb.servers[lb.roundrobin%len(lb.servers)]
	}
	lb.roundrobin++
	return server
}

// serveProxy method forwards the request to the selected backend server
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("forwarding request to address %q \n", targetServer.Address())
	targetServer.Serve(rw, r)
}

// handlErr function handles any errors, printing the error message and exiting the program
func handlErr(err error) {
	if err != nil {
		fmt.Println("error has occurred", err)
		os.Exit(1) // Exit if an error occurs
	}
}

func main() {
	// List of backend servers to forward traffic to
	servers := []Server{
		newsimplesever("http://localhost:8081"), // Backend server 1
		newsimplesever("http://localhost:8082"), // Backend server 3
		newsimplesever("http://localhost:8083"), // Backend server 2
	}
	lb := NewLoadBalancer("8080", servers) // Create a new load balancer listening on port 8000

	// Handle incoming requests and forward them to the backend servers
	handleRedirect := func(rw http.ResponseWriter, r *http.Request) {
		lb.serveProxy(rw, r) // Use the load balancer to serve the request
	}

	http.HandleFunc("/", handleRedirect)                        // Handle all requests to "/"
	fmt.Printf("serving requests at 'localhost:%s'\n", lb.port) // Log the port the load balancer is running on
	http.ListenAndServe(":8080", nil)                           // Start the HTTP server on the specified port

}
