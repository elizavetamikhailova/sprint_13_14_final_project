package api

import (
	"elizavetamikhailova/sprint_13_14_final_project/db"
	"encoding/json"
	"errors"
	"fmt"
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
		writeJSON(w, TaskResponse{Error: "Invalid JSON format"})
		return
	}

	if err := validateAndAdjustTask(&task); err != nil {
		writeJSON(w, TaskResponse{Error: err.Error()})
		return
	}

	fmt.Printf("will add to db date from task is %s \n", task.Date)
	id, err := db.AddTask(&task)
	if err != nil {
		writeJSON(w, TaskResponse{Error: "Failed to add task to database"})
		return
	}

	writeJSON(w, TaskResponse{ID: id})
}

func validateAndAdjustTask(task *db.Task) error {
	if task.Title == "" {
		return errors.New("task title is required")
	}

	now := time.Now()
	currentDate := now.Format(dateFormat)

	// Обработка специального значения "today"
	if task.Date == "today" {
		task.Date = currentDate
		return nil
	}

	// Если дата не указана, используем текущую дату
	if task.Date == "" {
		task.Date = currentDate
	}

	// Проверяем формат даты
	t, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return errors.New("invalid date format, expected YYYYMMDD")
	}

	// Если есть правило повторения, проверяем его и вычисляем следующую дату
	if task.Repeat != "" {
		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
		task.Date = next
	} else if afterNow(now, t) {
		// Если дата в прошлом и нет правила повторения, используем текущую дату
		task.Date = currentDate
	}

	return nil
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}
