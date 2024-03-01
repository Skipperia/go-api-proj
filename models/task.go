package models

type Task struct {
	ID   int
	Name string
}

var tasks = []Task{
	{ID: 1, Name: "Task One"},
	{ID: 2, Name: "Task Two"},
}

// GetAllTasks returns all tasks
func GetAllTasks() []Task {
	return tasks
}
