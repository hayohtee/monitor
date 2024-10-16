package main

import (
	"net/http"

	"github.com/hayohtee/monitor/internal/htmx"
)

type server struct {
	subscriberMessageBuffer int
	mux                     http.ServeMux
}

func NewServer() *server {
	s := server{
		subscriberMessageBuffer: 10,
	}
	s.mux.Handle("/", http.FileServerFS(htmx.PageFS))
	return &s
}
