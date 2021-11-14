package domain

import "time"

type User struct {
	Id        string    `json:"id" bson:"_id,omitempty" db:"id"`
	Login     string    `json:"login" bson:"login" db:"login"`
	Password  string    `json:"-" bson:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" db:"created_at"`
}

type CreateUserInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
