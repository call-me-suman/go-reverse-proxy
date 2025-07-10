package ratelimiter

import (
	"sync"
	"time"
	"net/http"
)

type Client struct{
	Tokens int
	LastSeen time.Time
}

type RateLimiter struct {
	MaxRequest int
	Window time.Duration
	clients map[string]*Client
	mu sync.Mutex
}

func NewRateLimiter(maxreq int,seconds int) *RateLimiter{

	r1 := &RateLimiter{
		MaxRequest:maxreq,
		Window:time.Second*time.Duration(seconds),
		clients:make(map[string]*Client),
	}
	go func(){
			ticker := time.NewTicker(r1.Window)
		
		for range ticker.C{
			r1.mu.Lock()
			for _, clients := range(r1.clients){

				clients.Tokens = r1.MaxRequest
				clients.LastSeen = time.Now()
			}
			r1.mu.Unlock()
		
		}
	}()

	return r1



}

func (r1 *RateLimiter)MiddleWare(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		ip:= r.RemoteAddr

		r1.mu.Lock()
		client , exists := r1.clients[ip]
		if !exists{
			client = &Client{
				Tokens:r1.MaxRequest -1,
				LastSeen : time.Now(),
			}

			r1.clients[ip] = client
		}else{
			if client.Tokens<=0{
				r1.mu.Unlock()
				http.Error(w, "Too many requests" , http.StatusTooManyRequests)
				return 
				
			}
			client.Tokens --
		}

		r1.mu.Unlock()
		next.ServeHTTP(w,r)
	})
}