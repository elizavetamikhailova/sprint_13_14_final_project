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

	// Получаем задачу из базы данных
	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, TaskResponse{Error: err.Error()})
		return
	}

	now := time.Now()

	if task.Repeat == "" {
		// Удаляем одноразовую задачу
		if err := db.DeleteTask(id); err != nil {
			writeJSON(w, TaskResponse{Error: "Failed to delete task"})
			return
		}
	} else {
		// Для периодической задачи вычисляем следующую дату
		nextDate, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeJSON(w, TaskResponse{Error: err.Error()})
			return
		}

		// Обновляем дату выполнения задачи
		if err := db.UpdateTaskDate(id, nextDate); err != nil {
			writeJSON(w, TaskResponse{Error: "Failed to update task date"})
			return
		}
	}

	writeJSON(w, TaskResponse{})
}