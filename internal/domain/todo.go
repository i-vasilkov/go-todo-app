package domain

import "time"

type Todo struct {
	Id        string    `json:"id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name,omitempty"`
	UserId    string    `json:"user_id" bson:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type UpdateTodoInput struct {
	Name string `json:"name" binding:"required"`
}

type CreateTodoInput struct {
	Name string `json:"name" binding:"required"`
}
