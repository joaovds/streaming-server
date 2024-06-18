package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joaovds/streaming-server/internal/handler"
	"github.com/joaovds/streaming-server/internal/infra/database"
)

func init() {
}

func main() {
	db, _ := database.SetupDatabase()
	defer db.Close()

	mux := http.NewServeMux()
	authHandler := handler.NewKeyHandler()

	mux.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("POST /auth", authHandler.Auth)

	port := "8080"
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
	log.Println(fmt.Sprintf("Starting server on http://localhost:%s", port))
}
