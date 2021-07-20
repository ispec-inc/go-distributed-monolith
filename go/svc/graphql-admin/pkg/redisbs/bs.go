package redisbs

import (
	"github.com/ispec-inc/monorepo/go/pkg/msgbs"
	"github.com/ispec-inc/monorepo/go/pkg/redis"
	"github.com/ispec-inc/monorepo/go/svc/admin/pkg/config"
)

var bs *msgbs.RedisMessageBus

func Init() (err error) {
	rcon, err := redis.New(redis.Config{
		Host: config.RedisMsgbs.Host,
		Port: config.RedisMsgbs.Port,
	})
	if err != nil {
		return err
	}

	pscon, err := redis.NewPubSub(redis.PubSubConfig{
		Host: config.RedisMsgbs.Host,
		Port: config.RedisMsgbs.Port,
	})
	if err != nil {
		return err
	}
	bs = msgbs.NewRedis(pscon, &rcon)

	return nil
}

func Get() *msgbs.RedisMessageBus {
	return bs
}
