package registry

import (
	"github.com/ispec-inc/monorepo/go/pkg/msgbs"
	"github.com/ispec-inc/monorepo/go/pkg/redis"
	"github.com/ispec-inc/monorepo/go/svc/article/pkg/config"
)

type MessageBus struct {
	msgbs msgbs.MessageBus
}

func NewMessageBus() (MessageBus, error) {

	rcon, err := redis.New(redis.Config{
		Host: config.RedisMsgbs.Host,
		Port: config.RedisMsgbs.Port,
	})
	if err != nil {
		return MessageBus{}, err
	}

	pscon, err := redis.NewPubSub(redis.PubSubConfig{
		Host: config.RedisMsgbs.Host,
		Port: config.RedisMsgbs.Port,
	})
	if err != nil {
		return MessageBus{}, err
	}

	bs := msgbs.NewRedis(pscon, &rcon)
	if err != nil {
		return MessageBus{}, err
	}

	return MessageBus{bs}, nil
}

func (m MessageBus) New() msgbs.MessageBus {
	return m.msgbs
}
