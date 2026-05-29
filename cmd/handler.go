package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func dataAccessHandler(w http.ResponseWriter, r *http.Request) {

	msg := Response{
		Message: "You are accessing data",
	}

	w.Header().Set("Content-Type", "application/json")

	js, err := json.MarshalIndent(msg, "", "\t") // dents things that are at depth level more than 0

	if err != nil {
		log.Fatal("Some error in parsing struct", err)
	}

	w.Write(js)

}
