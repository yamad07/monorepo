package msgbs

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type Subscriber interface {
	Do(redis.Message) error
}

type Router struct {
	Subscribers map[Event]Subscriber
}

func NewRouter() Router {
	return Router{
		Subscribers: map[Event]Subscriber{},
	}
}

func (r Router) Mount(router Router) {
	for evnt, subsc := range router.Subscribers {
		r.Subscribers[evnt] = subsc
	}
}

func (r Router) Subscribe(evnt Event, subsc Subscriber) {
	r.Subscribers[evnt] = subsc
}

func (r Router) Serve() {
	ps := NewRedis()
	for evnt, subsc := range r.Subscribers {
		ps.Subscribe(evnt)
		go func(subsc Subscriber) {
			for {
				switch v := ps.Receive().(type) {
				case redis.Message:
					err := subsc.Do(v)
					log.Println(err)
				case redis.Subscription:
					fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
				case error:
					log.Println(v)
				}
			}
		}(subsc)
	}
}
