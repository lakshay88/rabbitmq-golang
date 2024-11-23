package models

type Task struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}
