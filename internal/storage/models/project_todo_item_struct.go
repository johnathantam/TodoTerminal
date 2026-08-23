package models

type TodoItem struct {
	ID     string     `json:"id"`
	Title  string     `json:"title"`
	Status TodoStatus `json:"status"`
}
