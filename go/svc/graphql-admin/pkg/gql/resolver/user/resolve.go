package user

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/graphql-go/graphql"
	"github.com/ispec-inc/monorepo/go/pkg/applog"
	"github.com/ispec-inc/monorepo/go/pkg/msgbs"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/logger"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/model"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/redisbs"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/view"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

type Resolver struct {
	log applog.AppLog
	bs  msgbs.MessageBus
}

func NewResolver() Resolver {
	return Resolver{
		log: applog.New(logger.Get()),
		bs:  redisbs.Get(),
	}
}

func (r Resolver) List(params graphql.ResolveParams) (interface{}, error) {
	usrs := model.Users{}
	if err := usrs.Find(); err != nil {
		r.log.Error(params.Context, err)
		return nil, err
	}

	src := &ast.Source{
		Name: "./svc/graphql-admin/graphql/queries/user.graphql",
		Input: `
			type User{
				id: Int
				name: String
				email: String
			}
		`,
	}
	s, err := gqlparser.LoadSchema(src)
	if err != nil {
		r.log.Error(params.Context, err)
		return nil, err
	}
	spew.Dump(s.Types["User"].Fields[0].Name)

	if err := usrs.Find(); err != nil {
		r.log.Error(params.Context, err)
		return nil, err
	}

	as := view.NewUsers(&usrs)
	return as, nil
}
