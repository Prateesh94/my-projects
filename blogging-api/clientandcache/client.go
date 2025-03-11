package clientandcache

import (
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

var clients = make(map[string]*Client)
var mu sync.Mutex

type Client struct {
	limiter *rate.Limiter
}

// Adding and checking if the current client already exists
func getclient(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	if client, exists := clients[ip]; exists {
		fmt.Println("Client Found")
		return client.limiter
	}
	limit := rate.NewLimiter(rate.Limit(50), 1)
	clients[ip] = &Client{limit}
	fmt.Println("Added client")
	return limit
}

// middleware to handle multiple client request
func Limitmid(a http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		lim := getclient(ip)
		if !lim.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			fmt.Println("error generated")
			return
		}
		a.ServeHTTP(w, r)
	})
}
