package model

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Done        bool   `json:"done"`
}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var TaskDB map[string]Task

func init() {
	TaskDB = make(map[string]Task, 0)
}
