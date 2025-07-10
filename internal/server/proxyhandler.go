package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
	"strings"
	"net/url"
	
)


func NewProxy(target *url.URL) *httputil.ReverseProxy{
	proxy := httputil.NewSingleHostReverseProxy(target)
	return proxy
}


func ProxyHandler(proxy *httputil.ReverseProxy, url *url.URL,endpoint string) func(http.ResponseWriter , *http.Request){
	
	return func(w http.ResponseWriter , r *http.Request){
		
		
		fmt.Printf("Request received at %s at %s\n", r.URL, time.Now().UTC())

		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host",r.URL.Host)
		r.Host = url.Host

		path := r.URL.Path
		// fmt.Println(path,endpoint)
		r.URL.Path= strings.TrimLeft(path,endpoint)

		 fmt.Printf(" Redirecting request to %s at %s\n", r.URL, time.Now().UTC())


		proxy.ServeHTTP(w, r)
	}
}