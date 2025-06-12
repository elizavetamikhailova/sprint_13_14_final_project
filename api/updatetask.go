package api

import (
	"elizavetamikhailova/sprint_13_14_final_project/db"
	"encoding/json"
	"net/http"
)

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, TaskResponse{Error: "Invalid JSON format"}, http.StatusBadRequest)
		return
	}

	if task.ID == 0 {
		writeJSON(w, TaskResponse{Error: "Task ID is required"}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJSON(w, TaskResponse{Error: "Task title is required"}, http.StatusBadRequest)
		return
	}

	if err := validateAndAdjustTask(&task); err != nil {
		writeJSON(w, TaskResponse{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		writeJSON(w, TaskResponse{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	writeJSON(w, TaskResponse{ID: task.ID}, http.StatusOK)
}
