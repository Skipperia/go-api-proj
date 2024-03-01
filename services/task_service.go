package services

import "go-api-proj/models"

// GetTasks returns all tasks
func GetTasks() []models.Task {
	// In real scenario, add business logic here
	return models.GetAllTasks()
}
