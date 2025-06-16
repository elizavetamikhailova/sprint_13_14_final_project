package api

import (
	"elizavetamikhailova/sprint_13_14_final_project/db"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type TaskResponse struct {
	ID    int64  `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, TaskResponse{Error: "Invalid JSON format"}, http.StatusBadRequest)
		return
	}

	if err := validateAndAdjustTask(&task); err != nil {
		writeJSON(w, TaskResponse{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJSON(w, TaskResponse{Error: "Failed to add task to database"}, http.StatusBadRequest)
		return
	}

	writeJSON(w, TaskResponse{ID: id}, http.StatusOK)
}

func validateAndAdjustTask(task *db.Task) error {
	if task.Title == "" {
		return errors.New("task title is required")
	}

	now := time.Now()
	currentDate := now.Format(dateFormat)

	if task.Date == "" {
		task.Date = currentDate
	}

	t, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return errors.New("invalid date format, expected YYYYMMDD")
	}

	if task.Repeat != "" {
		if today(now, t) {
			task.Date = currentDate
		} else {
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = next
		}
		
	} else if afterNow(now, t) {
		task.Date = currentDate
	}

	return nil
}