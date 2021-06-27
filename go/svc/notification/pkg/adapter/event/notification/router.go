package notification

import (
	"github.com/ispec-inc/monorepo/go/pkg/msgbs"
	"github.com/ispec-inc/monorepo/go/svc/notification/pkg/registry"
)

func NewRouter(rgst registry.Registry) msgbs.Router {
	r := msgbs.NewRouter()

	subsc := NewSubscriber(rgst)

	r.Subscribe(msgbs.AddArticle, subsc.Notify)

	return r
}
