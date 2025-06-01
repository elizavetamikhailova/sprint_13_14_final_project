package api

import (
	"elizavetamikhailova/sprint_13_14_final_project/db"
	"encoding/json"
	"net/http"
)

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
    var task db.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        writeJSON(w, "Invalid JSON format")
        return
    }

    if task.ID == 0 {
        writeJSON(w, "Task ID is required")
        return
    }

    if task.Title == "" {
        writeJSON(w, "Task title is required")
        return
    }

    if err := validateAndAdjustTask(&task); err != nil {
        writeJSON(w, err.Error())
        return
    }

    if err := db.UpdateTask(&task); err != nil {
        writeJSON(w, err.Error())
        return
    }

    writeJSON(w, map[string]interface{}{})
}