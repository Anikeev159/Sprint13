// pkg/api/api.go
package api

import "net/http"

func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", taskHandler) // ← добавили
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task/done", doneTaskHandler)
}

// pkg/api/api.go
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodPut:
		updateTaskHandler(w, r)
	case http.MethodDelete:
		deleteTaskHandler(w, r) // ← новая строка
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
