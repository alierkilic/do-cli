package model

type NewTask struct {
	Task  string `json:"task"`
	Daily bool   `json:"daily"`
}

type Task struct {
	ID    int    `json:"id"`
	Task  string `json:"task"`
	Done  bool   `json:"done"`
	Daily bool   `json:"daily"`
}
