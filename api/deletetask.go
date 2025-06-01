package api

import (
	"elizavetamikhailova/sprint_13_14_final_project/db"
	"net/http"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSON(w, TaskResponse{Error: "Task ID is required"})
		return
	}


    if err := db.DeleteTask(id); err != nil {
        if err.Error() == "task not found" {
            writeJSON(w, TaskResponse{Error: err.Error()})
        } else {
            writeJSON(w, TaskResponse{Error: "Failed to delete task"})
        }
        return
    }

    writeJSON(w, TaskResponse{})
}