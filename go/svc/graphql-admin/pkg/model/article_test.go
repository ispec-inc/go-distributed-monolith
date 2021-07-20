package model

import (
	"testing"

	"github.com/ispec-inc/monorepo/go/pkg/infra/entity"
	"github.com/ispec-inc/monorepo/go/pkg/rdb"
	"github.com/ispec-inc/monorepo/go/pkg/testool"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestArticleModel_Find(t *testing.T) {
	type (
		give struct {
			id int64
		}
		want struct {
			id int64
		}
	)
	tests := []struct {
		name string
		give give
		want want
		err  bool
	}{
		{
			name: "success",
			give: give{id: int64(1)},
			want: want{id: int64(1)},
			err:  false,
		},
		{
			name: "not found",
			give: give{id: int64(2)},
			want: want{id: int64(0)},
			err:  true,
		},
	}

	db, cleanup := testool.PrepareDB(t, rdb.DBMSMySQL, "article_model_find", []interface{}{
		&entity.Article{ID: int64(1)},
	})
	database.Init(db)
	defer cleanup()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			atl := &Article{}
			err := atl.Find(tt.give.id)
			assert.Equal(t, tt.err, err != nil)
			assert.Equal(t, tt.want.id, atl.ID)
		})
	}
}

func TestArticleModel_Create(t *testing.T) {
	type (
		give struct {
			article *Article
		}
		want struct {
			changedCount int
		}
	)
	tests := []struct {
		name string
		give give
		want want
		err  bool
	}{
		{
			name: "success",
			give: give{article: NewArticle(1, "title", "body")},
			want: want{changedCount: 1},
			err:  false,
		},
		{
			name: "empty body",
			give: give{article: NewArticle(1, "title", "")},
			want: want{changedCount: 0},
			err:  true,
		},
	}

	db, cleanup := testool.PrepareDB(t, rdb.DBMSMySQL, "article_model_create", []interface{}{
		&entity.Article{ID: int64(1)},
	})
	database.Init(db)
	defer cleanup()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			bcnt := testool.CountRecord(t, db, entity.ArticleTableName)
			err := tt.give.article.Create()
			acnt := testool.CountRecord(t, db, entity.ArticleTableName)
			assert.Equal(t, tt.err, err != nil)
			assert.Equal(t, tt.want.changedCount, acnt-bcnt)
		})
	}
}

func TestArticleModel_Save(t *testing.T) {
	type (
		give struct {
			article *Article
		}
		want struct {
			title string
		}
	)
	tests := []struct {
		name string
		give give
		want want
		err  bool
	}{
		{
			name: "success",
			give: give{article: NewArticle(1, "new title", "body")},
			want: want{title: "new title"},
			err:  false,
		},
		{
			name: "empty title",
			give: give{article: NewArticle(1, "", "body")},
			want: want{title: ""},
			err:  true,
		},
	}

	db, cleanup := testool.PrepareDB(t, rdb.DBMSMySQL, "article_model_save", []interface{}{
		&entity.Article{ID: int64(1)},
	})
	database.Init(db)
	defer cleanup()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.give.article.Save()
			atl := &entity.Article{}
			db.First(atl, tt.give.article.ID)
			assert.Equal(t, tt.err, err != nil)
			assert.Equal(t, tt.want.title, atl.Title)
		})
	}
}

func TestArticleModel_Delete(t *testing.T) {
	type (
		give struct {
			id int64
		}
		want struct {
			changedCount int
		}
	)
	tests := []struct {
		name string
		give give
		want want
		err  bool
	}{
		{
			name: "success",
			give: give{id: int64(1)},
			want: want{changedCount: -1},
			err:  false,
		},
	}

	db, cleanup := testool.PrepareDB(t, rdb.DBMSMySQL, "article_model_delete", []interface{}{
		&entity.Article{ID: int64(1)},
	})
	database.Init(db)
	defer cleanup()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			atl := &Article{}
			atl.Find(tt.give.id)
			bcnt := testool.CountRecord(t, db, entity.ArticleTableName)
			err := atl.Delete()
			acnt := testool.CountRecord(t, db, entity.ArticleTableName)
			assert.Equal(t, tt.err, err != nil)
			assert.Equal(t, tt.want.changedCount, acnt-bcnt)
		})
	}
}

func TestArticlesModel_Find(t *testing.T) {
	type (
		want struct {
			count int
		}
	)
	tests := []struct {
		name string
		want want
		err  bool
	}{
		{
			name: "success",
			want: want{count: 1},
			err:  false,
		},
	}

	db, cleanup := testool.PrepareDB(t, rdb.DBMSMySQL, "articles_model_find", []interface{}{
		&entity.Article{ID: int64(1)},
	})
	database.Init(db)
	defer cleanup()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			atls := &Articles{}
			err := atls.Find()
			assert.Equal(t, tt.err, err != nil)
			assert.Equal(t, tt.want.count, len(*atls))
		})
	}
}
