package api

import (
	"encoding/json"
	"go-api-proj/models"
	"net/http"
)

// GetTasksHandler handles the "/tasks" route
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := models.GetAllTasks()
	json.NewEncoder(w).Encode(tasks)
}

// NewRouter creates a new HTTP router
func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/tasks", GetTasksHandler)
	return router
}
