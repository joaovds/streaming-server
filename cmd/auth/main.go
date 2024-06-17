package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("POST /auth", func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		defer body.Close()

		data, _ := io.ReadAll(body)
		fmt.Println("Data: ", string(data))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	port := "8080"
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
	log.Println(fmt.Sprintf("Starting server on http://localhost:%s", port))
}
