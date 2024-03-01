package repository

import "go-api-proj/models"

// GetTasks returns all tasks from the data source
func GetTasks() []models.Task {
	// In real scenario, this would interact with a database
	return models.GetAllTasks()
}
