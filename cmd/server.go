package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/hayohtee/monitor/static"
)

type server struct {
	subscriberMessageBuffer int
	mux                     http.ServeMux
	subscribers             map[*subscriber]struct{}
	mu                      sync.Mutex
}

func NewServer() *server {
	s := server{
		subscriberMessageBuffer: 10,
		subscribers:             make(map[*subscriber]struct{}),
	}

	s.mux.Handle("/", http.FileServerFS(static.Files))
	s.mux.HandleFunc("/ws", s.subscribeHandler)

	return &s
}

func (s *server) subscribe(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var c *websocket.Conn
	sub := &subscriber{
		msgs: make(chan []byte, s.subscriberMessageBuffer),
	}
	s.addSubscriber(sub)

	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		return err
	}
	defer c.CloseNow()

	ctx = c.CloseRead(ctx)
	for {
		select {
		case msg := <-sub.msgs:
			ctx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()
			err := c.Write(ctx, websocket.MessageText, msg)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *server) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	err := s.subscribe(r.Context(), w, r)
	if err != nil {
		log.Println(err)
		return
	}

}

func (s *server) addSubscriber(sub *subscriber) {
	s.mu.Lock()
	s.subscribers[sub] = struct{}{}
	s.mu.Unlock()
	log.Println(sub, "subscriber added")
}

func (s *server) broadcast(msg []byte) {
	s.mu.Lock()
	for sub := range s.subscribers {
		sub.msgs <- msg
	}
	s.mu.Unlock()
}
