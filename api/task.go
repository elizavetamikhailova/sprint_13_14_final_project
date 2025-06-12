package api

import (
	"elizavetamikhailova/sprint_13_14_final_project/db"
	"net/http"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
		writeJSON(w, TaskResponse{Error: "Task ID is required"}, http.StatusBadRequest)
        return
    }

    task, err := db.GetTask(id)
    if err != nil {
        writeJSON(w, TaskResponse{Error: err.Error()}, http.StatusBadRequest)
        return
    }

    writeJSON(w, task, http.StatusOK)
}