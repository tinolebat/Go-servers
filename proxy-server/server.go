package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var servercount = 0

const (
	SERVER1 = "https://www.google.com"
	SERVER2 = "https://www.facebook.com"
	SERVER3 = "https://www.yahoo.com"
	PORT    = "1338"
)

func loadbalancer(res http.ResponseWriter, req *http.Request) {

	//Get address of done backend server
	url := getProxyURL()

	// log the request
	logrequestPayload(url)

	// Forward request to original request
	serveReverseProxy(url, res, req)
}

func getProxyURL() string {
	var servers = []string{SERVER1, SERVER2, SERVER3}
	server := servers[servercount]
	servercount++

	if servercount >= len(servers) {
		servercount = 0
	}
	return server
}

func logrequestPayload(proxyurl string) {
	fmt.Printf("Proxy URL: %s\n", proxyurl)
}

func serveReverseProxy(targeturl string, res http.ResponseWriter, req *http.Request) {

	//parse the url
	url, _ := url.Parse(targeturl)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.ServeHTTP(res, req)

}

func tvserver(res http.ResponseWriter, req *http.Request) {

	fmt.Println("TV SERVER")

}

func Start() {
	http.HandleFunc("/", loadbalancer)
	http.HandleFunc("/tv", tvserver)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
	fmt.Println("Server Started...")
}

func main() {
	//Start Server
	fmt.Println("Load balancing proxy server started")
	Start()

}
