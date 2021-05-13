package msgbs

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
)

type RedisPubSub struct {
	Conn      *redis.PubSubConn
	RedisConn *redis.Conn
}

var redisHost string

func Init(host string, port string) {
	redisHost = host + ":" + port
}

func NewRedis() *RedisPubSub {
	redisConn, err := redis.Dial("tcp", redisHost)
	if err != nil {
		panic(err)
	}
	redisPubSub := &RedisPubSub{}
	redisPubSub.Conn = &redis.PubSubConn{Conn: redisConn}
	redisPubSub.RedisConn = &redisConn
	return redisPubSub
}

func (r *RedisPubSub) Publish(evnt Event, msg Message) error {
	j, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	redisConn := *r.RedisConn
	_, err = redisConn.Do("PUBLISH", evnt, string(j))

	return err
}

func (r *RedisPubSub) Subscribe(evnt Event) {
	r.Conn.Subscribe(evnt)
}

func (r *RedisPubSub) Receive() interface{} {
	return r.Conn.Receive()
}
