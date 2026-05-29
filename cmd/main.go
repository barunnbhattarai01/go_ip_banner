package main

import (
	"log"
	"net/http"
)

func main() {
	// simple http server with one endpoint

	http.HandleFunc("/", rateLimitAndIpBan(dataAccessHandler))
	err := http.ListenAndServe(":8080", nil) // nil means defualt we can pass other too like chi gorillamux router as well
	if err != nil {
		log.Fatal("Some error while making a http server", err)
	}

}
