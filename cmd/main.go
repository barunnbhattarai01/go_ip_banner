package main

import (
	"log"
	"net/http"
)

func main() {
	// simple http server with one endpoint

	http.HandleFunc("/", rateLimitAndIpBan(dataAccessHandler))
	log.Println("Server is listening on port 8080 on all interfaces")
	err := http.ListenAndServe("0.0.0.0:8080", nil) // nil means defualt we can pass other too like chi gorillamux router as well
	if err != nil {
		log.Fatal("Some error while making a http server", err)
	}

}
