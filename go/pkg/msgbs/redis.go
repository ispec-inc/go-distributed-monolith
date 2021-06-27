package msgbs

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
)

type RedisMessageBus struct {
	PubSubConn *redis.PubSubConn
	Conn       *redis.Conn
}

func NewRedis(pscon *redis.PubSubConn, conn *redis.Conn) *RedisMessageBus {
	return &RedisMessageBus{
		PubSubConn: pscon,
		Conn:       conn,
	}
}

func (r RedisMessageBus) Publish(evnt Event, msg Message) error {
	j, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	rcon := *r.Conn
	_, err = rcon.Do("PUBLISH", evnt, string(j))

	return err
}

func (r RedisMessageBus) Subscribe(evnt Event) {
	r.PubSubConn.Subscribe(evnt)
}

func (r RedisMessageBus) Unsubscribe(evnt Event) {
	r.PubSubConn.Unsubscribe(evnt)
}

func (r RedisMessageBus) Receive() interface{} {
	return r.PubSubConn.Receive()
}
