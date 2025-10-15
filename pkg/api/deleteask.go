// pkg/api/deletetask.go
package api

import (
<<<<<<< HEAD
=======
	"encoding/json"
>>>>>>> 246564f55ef46378affc64b1c8dc75b4b0d31cd3
	"net/http"

	"go-final-project/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
<<<<<<< HEAD
		WriteJSON(w, map[string]string{"error": "id is required"}, http.StatusBadRequest)
=======
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "id is required"})
>>>>>>> 246564f55ef46378affc64b1c8dc75b4b0d31cd3
		return
	}

	if err := db.DeleteTask(id); err != nil {
<<<<<<< HEAD
		WriteJSON(w, map[string]string{"error": "task not found"}, http.StatusNotFound)
		return
	}

	WriteJSON(w, map[string]interface{}{}, http.StatusOK)
=======
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{})
>>>>>>> 246564f55ef46378affc64b1c8dc75b4b0d31cd3
}
