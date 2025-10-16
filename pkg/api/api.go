// pkg/api/api.go
package api

import (
	"encoding/json"
	"net/http"
)

func Init() {
	http.HandleFunc("GET /api/tasks", tasksHandler)

	http.HandleFunc("POST /api/task/done", doneTaskHandler)

	http.HandleFunc("GET /api/task", getTaskHandler)

	http.HandleFunc("POST /api/task", addTaskHandler)

	http.HandleFunc("PUT /api/task", updateTaskHandler)

	http.HandleFunc("DELETE /api/task", deleteTaskHandler)

	http.HandleFunc("/api/nextdate", nextDateHandler)
}

func WriteJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
