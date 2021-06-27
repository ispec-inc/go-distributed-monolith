package event

import (
	"github.com/ispec-inc/monorepo/go/pkg/applog"
	"github.com/ispec-inc/monorepo/go/pkg/msgbs"
	"github.com/ispec-inc/monorepo/go/svc/notification/pkg/adapter/event/notification"
	"github.com/ispec-inc/monorepo/go/svc/notification/pkg/registry"
)

func NewSubscriber() (msgbs.Subscriber, func() error, error) {
	lgr, lgrCleanup, err := registry.NewLogger()
	if err != nil {
		return msgbs.Subscriber{}, nil, err
	}
	cleanup := func() error {
		lgrCleanup()
		return nil
	}

	rgst := registry.NewRegistry(lgr)
	mbs, err := registry.NewMessageBus()
	if err != nil {
		return msgbs.Subscriber{}, nil, err
	}

	r := msgbs.NewSubscriber(
		mbs.New(),
		applog.New(lgr.New()),
	)

	nr := notification.NewRouter(rgst)

	r.Mount(nr)
	return r, cleanup, nil
}
