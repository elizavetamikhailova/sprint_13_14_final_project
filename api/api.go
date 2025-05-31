package api

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func RunServer() {
	r := chi.NewRouter()

	webDir := "./web"
	fileServer := http.FileServer(http.Dir(webDir))
	r.Handle("/*", fileServer)

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	r.Get("/api/nextdate", NextDateHandler)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Произошла ошибка при запуске сервера %v", err)
	}

	//http.HandleFunc("/api/nextdate", NextDateHandler)
}
