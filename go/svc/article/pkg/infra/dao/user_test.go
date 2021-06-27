package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ispec-inc/monorepo/go/pkg/infra/entity"
	"github.com/ispec-inc/monorepo/go/pkg/rdb"
	"github.com/ispec-inc/monorepo/go/pkg/testool"
	"github.com/ispec-inc/monorepo/go/svc/article/pkg/domain/model"
)

func TestUserDao_Get(t *testing.T) {
	db, cleanup := testool.PrepareDB(t, rdb.DBMSMySQL, "user_dao_get", []interface{}{
		&entity.User{ID: int64(1), Email: "email"},
	})
	defer cleanup()
	d := NewUser(db)

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
		},
		{
			name: "not found",
			give: give{id: int64(2)},
			want: want{id: int64(0)},
			err:  true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, aerr := d.Get(tt.give.id)
			if tt.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
				assert.Exactly(t, tt.want.id, got.ID)
			}
		})
	}
}

func TestUserDao_List(t *testing.T) {
	db, cleanup := testool.PrepareDB(t, rdb.DBMSMySQL, "user_dao_find_by_user_id", []interface{}{
		&entity.User{ID: int64(1), Email: "email-1"},
		&entity.User{ID: int64(2), Email: "email-2"},
	})
	defer cleanup()
	d := NewUser(db)

	type (
		give struct {
			ids []int64
		}
		want struct {
			count int
		}
	)
	tests := []struct {
		name string
		give give
		want want
		err  bool
	}{
		{
			name: "success without ids",
			give: give{ids: nil},
			want: want{count: 2},
			err:  false,
		},
		{
			name: "success with ids",
			give: give{ids: []int64{1}},
			want: want{count: 1},
			err:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, aerr := d.List(tt.give.ids)
			if tt.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
				assert.Exactly(t, tt.want.count, len(got))
			}
		})
	}
}

func TestUserDao_Create(t *testing.T) {
	db, cleanup := testool.PrepareDB(t, rdb.DBMSMySQL, "user_dao_create", []interface{}{
		&entity.User{ID: int64(1), Email: "old"},
	})
	defer cleanup()
	d := NewUser(db)

	type (
		give struct {
			user *model.User
		}
		want struct {
			createdCount int
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
			give: give{user: &model.User{Email: "new"}},
			want: want{createdCount: 1},
			err:  false,
		},
		{
			name: "fail for already existed",
			give: give{user: &model.User{Email: "old"}},
			want: want{createdCount: 0},
			err:  true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			bcnt := testool.CountRecord(t, db, "users")
			aerr := d.Create(tt.give.user)
			acnt := testool.CountRecord(t, db, "users")

			if tt.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
				assert.Exactly(t, tt.want.createdCount, acnt-bcnt)
			}
		})
	}
}
