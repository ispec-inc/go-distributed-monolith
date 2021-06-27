package msgbs

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/ispec-inc/monorepo/go/pkg/applog"
)

type SubscribeFunc func(context.Context, redis.Message) error

type Subscriber struct {
	MessageBus MessageBus
	Router     Router
	AppLog     applog.AppLog
}

func NewSubscriber(msgbs MessageBus, algr applog.AppLog) Subscriber {
	return Subscriber{
		MessageBus: msgbs,
		AppLog:     algr,
		Router:     Router{},
	}
}

func (s Subscriber) Subscribe(evnt Event, subsc SubscribeFunc) {
	s.Router.Subscribe(evnt, subsc)
}

func (s Subscriber) Mount(rtr Router) {
	for evnt, sfuncs := range rtr {
		s.Router[evnt] = append(s.Router[evnt], sfuncs...)
	}
}

func (s Subscriber) Do(ctx context.Context) {
	for evnt, sfuncs := range s.Router {
		s.MessageBus.Subscribe(evnt)
		for _, sfunc := range sfuncs {
			go func(sfunc SubscribeFunc) {
				for {
					switch v := s.MessageBus.Receive().(type) {
					case redis.Message:
						err := sfunc(ctx, v)
						if err != nil {
							s.AppLog.Error(ctx, err)
						}
					case redis.Subscription:
						fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
					case error:
						s.AppLog.Error(ctx, v)
					}
				}
			}(sfunc)
		}
	}
}
