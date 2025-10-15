// pkg/api/addtask.go
package api

import (
	"encoding/json"
	"net/http"

<<<<<<< HEAD
=======
	"go-final-project/pkg/api"
>>>>>>> 246564f55ef46378affc64b1c8dc75b4b0d31cd3
	"go-final-project/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
<<<<<<< HEAD
		WriteJSON(w, map[string]string{"error": "invalid JSON"}, http.StatusBadRequest)
=======
		api.WriteJSON(w, map[string]string{"error": "invalid JSON"}, http.StatusBadRequest)
>>>>>>> 246564f55ef46378affc64b1c8dc75b4b0d31cd3
		return
	}

	if task.Title == "" {
<<<<<<< HEAD
		WriteJSON(w, map[string]string{"error": "title is required"}, http.StatusBadRequest)
		return
	}

	if err := checkDate(&task); err != nil { // ← без параметра now
		WriteJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
=======
		api.WriteJSON(w, map[string]string{"error": "title is required"}, http.StatusBadRequest)
		return
	}

	if err := checkDate(&task); err != nil {
		api.WriteJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
>>>>>>> 246564f55ef46378affc64b1c8dc75b4b0d31cd3
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
<<<<<<< HEAD
		WriteJSON(w, map[string]string{"error": "database error"}, http.StatusInternalServerError)
		return
	}

	WriteJSON(w, map[string]int64{"id": id}, http.StatusCreated)
=======
		api.WriteJSON(w, map[string]string{"error": "database error"}, http.StatusInternalServerError)
		return
	}

	api.WriteJSON(w, map[string]int64{"id": id}, http.StatusCreated)
>>>>>>> 246564f55ef46378affc64b1c8dc75b4b0d31cd3
}
