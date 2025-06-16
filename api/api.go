package api

import (
	"encoding/json"
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
	r.Post("/api/task", auth(AddTaskHandler))
	r.Get("/api/tasks", auth(TasksHandler))
	r.Get("/api/task", TaskHandler)
	r.Put("/api/task", UpdateTaskHandler)
	r.Post("/api/task/done", auth(TaskDoneHandler))
	r.Delete("/api/task", DeleteTaskHandler)
	r.Post("/api/signin", SignInHandler)


	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Произошла ошибка при запуске сервера %v", err)
	}
}

func writeJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
