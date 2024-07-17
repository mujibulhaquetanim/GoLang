package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	}))
	router.HandleFunc("/first/{id}", func(w http.ResponseWriter, r *http.Request) {
		id:= r.PathValue("id")
		w.Write([]byte("received req for id " + id))
	})

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	log.Println("Listening and serving HTTP on http://127.0.0.1:8080")

	server.ListenAndServe()
}
