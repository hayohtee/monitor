package main

import (
	"net/http"

	"github.com/hayohtee/monitor/static"
)

type server struct {
	subscriberMessageBuffer int
	mux                     http.ServeMux
}

func NewServer() *server {
	s := server{
		subscriberMessageBuffer: 10,
	}
	s.mux.Handle("/", http.FileServerFS(static.Files))
	return &s
}
