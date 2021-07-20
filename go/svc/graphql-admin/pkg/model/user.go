package model

import (
	"errors"

	"github.com/ispec-inc/monorepo/go/pkg/apperror"
	"github.com/ispec-inc/monorepo/go/pkg/infra/entity"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/database"
	"go.uber.org/multierr"
)

type (
	User struct {
		entity.User
	}
	Users []User
)

func NewUser(name string, desc string, email string) *User {
	usr := &User{}
	usr.Name = name
	usr.Description = desc
	usr.Email = email
	return usr
}

func (u *User) validate() error {
	var err error
	if u.Name == "" {
		err = multierr.Append(err, errors.New("user must have name"))
	}

	if u.Description == "" {
		err = multierr.Append(err, errors.New("user must have name"))
	}

	if u.Email == "" {
		err = multierr.Append(err, errors.New("user must have name"))
	}

	return nil
}

func (a *User) Find(id int64) error {
	return apperror.NewGormFind(
		database.Get().First(a, id).Error,
		entity.UserTableName,
	)
}

func (a *User) Create() error {
	if err := a.validate(); err != nil {
		return err
	}
	return apperror.NewGormCreate(
		database.Get().Create(a).Error,
		entity.UserTableName,
	)
}

func (a *User) Save() error {
	if err := a.validate(); err != nil {
		return err
	}
	return apperror.NewGormSave(
		database.Get().Save(a).Error,
		entity.UserTableName,
	)
}

func (a *User) Delete() error {
	return apperror.NewGormDelete(
		database.Get().Delete(a).Error,
		entity.UserTableName,
	)
}

func (a *Users) Find() error {
	return apperror.NewGormFind(
		database.Get().Find(a).Error,
		entity.UserTableName,
	)
}
