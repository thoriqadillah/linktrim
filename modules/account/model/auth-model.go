package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserCreate struct {
		Name     string `json:"name" validate:"required,alpha,min=2,max=32"`
		Email    string `json:"email" validate:"required,email,min=2,max=32"`
		Password string `json:"password" validate:"required,alhpanum,min=6,max=16"`
	}

	UserLogin struct {
		Email    string `json:"email" validate:"required,email,min=2,max=32"`
		Password string `json:"password" validate:"required,alhpanum,min=6,max=16"`
	}

	User struct {
		ID         uuid.UUID `json:"id" redis:"id"`
		Name       string    `json:"name" redis:"name"`
		Email      string    `json:"email" redis:"email"`
		Password   string    `json:"-" redis:"-"`
		ProfilePic string    `json:"profile_pic" redis:"profile_pic"`
		CreatedAt  time.Time `json:"created_at" redis:"created_at"`
		UpdatedAt  time.Time `json:"updated_at" redis:"updated_at"`
	}
)
