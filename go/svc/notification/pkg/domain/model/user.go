package model

import (
	"time"

	"github.com/ispec-inc/monorepo/go/pkg/infra/entity"
)

type User struct {
	ID          int64
	Name        string
	Description string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewUserFromEntity(e *entity.User) *User {
	return &User{
		CreatedAt:   e.CreatedAt,
		Description: e.Description,
		Email:       e.Email,
		ID:          e.ID,
		Name:        e.Name,
		UpdatedAt:   e.UpdatedAt,
	}
}
