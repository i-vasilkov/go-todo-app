package domain

import "time"

type Todo struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateTodoInput struct {
}

type CreateTodoInput struct {
}
