package msgbs

import (
	"context"
	"sync"
)

type SubscribeServer struct {
	subscribers []Subscriber
	done        chan interface{}
	wg          *sync.WaitGroup
}

func NewSubscribeServer() SubscribeServer {
	return SubscribeServer{
		done: make(chan interface{}),
		wg:   &sync.WaitGroup{},
	}
}

func (s SubscribeServer) Serve(ctx context.Context) {
	for _, subsc := range s.subscribers {
		subsc.Do(ctx, s.done, s.wg)
	}
}

func (s *SubscribeServer) Mount(subsc Subscriber) {
	s.subscribers = append(s.subscribers, subsc)
}

func (s SubscribeServer) Shutdown(ctx context.Context) {
	for _, subsc := range s.subscribers {
		for evnt := range subsc.Router {
			subsc.MessageBus.Unsubscribe(evnt)
		}
	}

	go func() {
		s.done <- true
	}()

	s.wg.Wait()
}
