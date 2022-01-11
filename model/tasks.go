package model

type NewTask struct {
	Task string `json:"task"`
}

type Task struct {
	ID   string `json:"id"`
	Task string `json:"task"`
	Done int    `json:"done"`
}
