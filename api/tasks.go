package api

import (
	"elizavetamikhailova/sprint_13_14_final_project/db"
	"net/http"
)

type TasksResponse struct {
	Tasks []*db.Task `json:"tasks"`
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {

	tasks, err := db.GetTasks(50)
	if err != nil {
		writeJSON(w, TaskResponse{Error: "Failed to get tasks"})
		return
	}

	writeJSON(w, TasksResponse{Tasks: tasks})
}
