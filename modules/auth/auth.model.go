package auth

import (
	"time"

	"github.com/google/uuid"
)

type (
	userCreate struct {
		Name     string `json:"name" validate:"required,alpha,min=2,max=32"`
		Email    string `json:"email" validate:"required,email,min=2,max=32"`
		Password string `json:"password" validate:"required,alhpanum,min=6,max=16"`
	}

	userLogin struct {
		Email    string `json:"email" validate:"required,email,min=2,max=32"`
		Password string `json:"password" validate:"required,alhpanum,min=6,max=16"`
	}

	User struct {
		ID         uuid.UUID `json:"id"`
		Name       string    `json:"name"`
		Email      string    `json:"email"`
		Password   string    `json:"-"`
		ProfilePic string    `json:"profile_pic"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
)
