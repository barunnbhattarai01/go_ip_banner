package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func rateLimitAndIpBan(next http.HandlerFunc) http.HandlerFunc {

	// at the time of initilizatio of this middleware run a go routine for clean up
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var mu sync.Mutex

	clients := make(map[string]*client)
	bannedClients := make(map[string]*client)

	go func() {

		//runs as long as the http server runs
		for {

			//so it runs every minute
			time.Sleep(time.Minute)

			mu.Lock()

			for ip, client := range clients {
				if time.Since(client.lastSeen) >= time.Minute*3 {
					delete(clients, ip)
				}

			}

			for ip, client := range bannedClients {
				if time.Since(client.lastSeen) > time.Minute*1440 {
					delete(clients, ip)

				}

			}

			mu.Unlock()
		}

	}()

	return func(w http.ResponseWriter, r *http.Request) {

		// add the clinet ip to the map if now found

		// extract the client's ip address form the request.
		// the ignored value here is the client empheral port which the os assigns
		ip, port, err := net.SplitHostPort(r.RemoteAddr)

		if err != nil {

			w.Write([]byte("Some server error"))
		}

		mu.Lock()

		if _, found := bannedClients[ip]; found {
			msg := fmt.Sprintf("your ip has been banned for 1 day for suspicious activity your public ip address is %s and empheral port is %s", ip, port)
			w.Write([]byte(msg))
			mu.Unlock()
			return
		}

		if _, found := clients[ip]; !found {
			clients[ip] = &client{
				limiter: rate.NewLimiter(2, 6),
			}
		}

		// update list seen time

		clients[ip].lastSeen = time.Now()

		if !clients[ip].limiter.Allow() {
			bannedClients[ip] = &client{
				limiter:  rate.NewLimiter(2, 6),
				lastSeen: time.Now(),
			}
			msg := fmt.Sprintf(" too many request your ip has been banned for 1 day for suspicious activity your public ip address is %s and empheral port is %s", ip, port)
			w.Write([]byte(msg))

			mu.Unlock()
			return
		}

		mu.Unlock()
		// calls the next request in the chain
		next.ServeHTTP(w, r)

	}

}
