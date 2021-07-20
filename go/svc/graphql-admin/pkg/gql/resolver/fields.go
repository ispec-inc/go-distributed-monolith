package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/gql/resolver/article"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/gql/resolver/user"
)

type QueryResolverRegistryImpl struct {
}

func NewQueryResolver() QueryResolverRegistryImpl {
	return QueryResolverRegistryImpl{}
}

func (q QueryResolverRegistryImpl) Users() func(params graphql.ResolveParams) (interface{}, error) {
	return user.NewResolver().List
}

func (q QueryResolverRegistryImpl) Articles() func(params graphql.ResolveParams) (interface{}, error) {
	return article.NewResolver().List
}
