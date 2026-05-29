package main

import (
	"log"
	"net/http"
)

func dataAccessHandler(w http.ResponseWriter, r *http.Request) {

	err := writeJson(w, envelope{"message": "You are accessing data"}, nil, 200)
	if err != nil {

		log.Fatal("some error occurred", err)
	}

}
