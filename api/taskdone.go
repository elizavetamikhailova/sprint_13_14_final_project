package api

import (
	"elizavetamikhailova/sprint_13_14_final_project/db"
	"net/http"
	"time"
)

func TaskDoneHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSON(w, TaskResponse{Error: "Task ID is required"})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, TaskResponse{Error: err.Error()})
		return
	}

	now := time.Now()

	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			writeJSON(w, TaskResponse{Error: "Failed to delete task"})
			return
		}
	} else {
		nextDate, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeJSON(w, TaskResponse{Error: err.Error()})
			return
		}

		if err := db.UpdateTaskDate(id, nextDate); err != nil {
			writeJSON(w, TaskResponse{Error: "Failed to update task date"})
			return
		}
	}

	writeJSON(w, TaskResponse{})
}