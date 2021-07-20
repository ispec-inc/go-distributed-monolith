package view

import (
	"time"

	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/model"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(m *model.User) User {
	return User{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func NewUsers(ms *model.Users) []User {
	vs := make([]User, len(*ms))
	for i, m := range *ms {
		vs[i] = NewUser(&m)
	}
	return vs
}
