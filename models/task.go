package models

import "time"

type Task struct {
	ID          int64 `json:"id"`
	Description string `json:"description"`
	Complete    bool `json:"complete"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
	User User `json:"user"`
}
