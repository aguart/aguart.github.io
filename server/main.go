package main

import (
	"log"
	"net/http"
)

func main() {
	s := &http.Server{
		Addr: ":8081",
	}

	fs := http.FileServer(http.Dir("../aguart.github.io"))
	http.Handle("/", fs)
	log.Println("Listening on :8081")
	log.Fatal(s.ListenAndServe())
}
