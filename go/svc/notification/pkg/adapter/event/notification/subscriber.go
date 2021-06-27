package notification

import (
	"context"
	"encoding/json"

	"github.com/gomodule/redigo/redis"
	"github.com/ispec-inc/monorepo/go/pkg/msgbs"
	"github.com/ispec-inc/monorepo/go/svc/notification/pkg/registry"
	"github.com/ispec-inc/monorepo/go/svc/notification/src/v1/notification"
)

type Subscriber struct {
	usecase notification.Usecase
}

func NewSubscriber(rgst registry.Registry) Subscriber {
	usecase := notification.NewUsecase(rgst)
	return Subscriber{usecase}
}

func (s Subscriber) Notify(ctx context.Context, msg redis.Message) error {
	var m msgbs.Article
	err := json.Unmarshal(msg.Data, &m)
	if err != nil {
		return err
	}

	ipt := notification.Input{
		Title: m.Title,
	}
	err = s.usecase.Send(ctx, &ipt)
	if err != nil {
		return err
	}

	return nil
}
