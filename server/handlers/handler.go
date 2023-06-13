package handlers

import (
	"github.com/RemcoE33/share-secret/store"
)

type Server struct {
	Storer store.Storer
}

func NewServer(store store.Storer) *Server {
	return &Server{
		Storer: store,
	}
}
