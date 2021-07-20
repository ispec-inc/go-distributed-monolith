package view

import (
	"time"

	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/model"
)

type Article struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewArticle(m *model.Article) Article {
	return Article{
		ID:        m.ID,
		UserID:    m.UserID,
		Title:     m.Title,
		Body:      m.Body,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func NewArticles(ms *model.Articles) []Article {
	vs := make([]Article, len(*ms))
	for i, m := range *ms {
		vs[i] = NewArticle(&m)
	}
	return vs
}
