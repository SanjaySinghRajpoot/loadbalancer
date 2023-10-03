package main

import (
	"net/http/httputil"
	"net/url"
)

type Server interface {

	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

type serverElement struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *serverElement {

	serverUrl, err := url.Parse(addr)

	return &serverElement{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl)
	}
}

type LoadBalancer struct {
	port string
	roundRobinCount int 
	servers []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
      return &LoadBalancer{
		roundRobinCount: 0,
		port: port,
		servers: servers
	  }
}


func handleError(err error)
{
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

// this is a method
func (s *serverElement) Address() string {return s.addr}

func (s *simpleServer)IsAlive() bool {return true}

func (lb *LoadBalancer) getNextAvailableServer() Server{
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]

	for !server.IsAlive() {
       lb.roundRobinCount++
	   server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}

	lb.roundRobinCount++
}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()

	fmt.Printf("forwarding request to address %q \n", targetServer.Address())

	targetServer.Serve(rw, req)
}

func main() {

	servers := []Server{
        newSimpleServer("https://www.facebook.com"),
        newSimpleServer("https://www.bing.com"),
        newSimpleServer("https://www.duckduckgo.com")		
	}

	lb := NewLoadBalancer("8000", servers)

	handleRedirect := func(rw http.ResponseWriter, req *http.Request){
       lb.serveProxy(rw, req)
	}
}