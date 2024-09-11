package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

type simpleserver struct {
	addr  string
	proxy httputil.ReverseProxy
}

func newsimplesever(addr string) *simpleserver {
	serverurl, err := url.Parse(addr)
	handlErr(err)
	return &simpleserver{
		addr:  addr,
		proxy: *httputil.NewSingleHostReverseProxy(serverurl),
	}
}
type LoadBalancer struct {
	port       string
	roundrobin int
	servers    []Server
}

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
	return true
}

func (s *simpleserver) Serve(rw http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(rw,r)
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	server:=lb.servers[lb.roundrobin%len(lb.servers)]
	for !server.IsAlive(){
		lb.roundrobin++
		server=lb.servers[lb.roundrobin%len(lb.servers)]
	}
	lb.roundrobin++
	return server
}
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, r *http.Request) {
	targetServer:=lb.getNextAvailableServer()
	fmt.Printf("forwading request to address %q \n", targetServer.Address())
	targetServer.Serve(rw ,r)
}

func handlErr(err error) {
	if err != nil {
		fmt.Println("error has occured", err)
		os.Exit(1) //os package provides a platform-independent interface to operating system functionality.
	}
}

func main() {
	servers := []Server{
		newsimplesever("https://facebook.com"),
		//newsimplesever("https://www.github.com"),
		newsimplesever("https://www.apple.com"),
	}
	lb := NewLoadBalancer("8000", servers)
	handleRedirect := func(rw http.ResponseWriter, r *http.Request) {
		lb.serveProxy(rw, r)
	}
	http.HandleFunc("/", handleRedirect)
	fmt.Printf("serving request at 'localhost:%s'", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
