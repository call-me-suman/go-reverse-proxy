package server

import (
	"net/http"
	"reverse-proxy/internal/configs"
	"fmt"
	"net/url"
	"reverse-proxy/internal/ratelimiter"
	"reverse-proxy/internal/auth"
)


func Run() error{

	config,err := configs.NewConfiguration()

	if err!= nil{
		return fmt.Errorf("Error : Cant Load the configurations")
	}

	r1:=ratelimiter.NewRateLimiter(config.Rate.Reqs,config.Rate.Persec) //5 request per 10 seconds

	mux := http.NewServeMux()

	mux.HandleFunc("/ping",ping)
	
	for _, res := range config.Resources{
		url, _ := url.Parse(res.Destination_url)
		proxy := NewProxy(url)
		proxyhandler:=ProxyHandler(proxy, url,res.Endpoint)
		mux.Handle(res.Endpoint ,auth.JWTAuthMiddleware("asdnaskdasiiyaebdbasgyyaefwe54641ef")(r1.MiddleWare(http.HandlerFunc(proxyhandler)) ))

	}

	if err := http.ListenAndServe(config.Server.Host +":"+config.Server.Listen_port,mux);err != nil{
		fmt.Errorf("Error in creating the server: %v",err)
	}


	return nil



}