package model

type NewTask struct {
	Task string `json:"task"`
}

type Task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done int    `json:"done"`
}
