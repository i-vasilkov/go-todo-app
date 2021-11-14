package domain

import "time"

type Task struct {
	Id        string    `json:"id" bson:"_id,omitempty" db:"id"`
	Name      string    `json:"name" bson:"name,omitempty" db:"name"`
	UserId    string    `json:"user_id" bson:"user_id,omitempty" db:"user_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at" db:"updated_at"`
}

type UpdateTaskInput struct {
	Name string `json:"name" binding:"required"`
}

type CreateTaskInput struct {
	Name string `json:"name" binding:"required"`
}
