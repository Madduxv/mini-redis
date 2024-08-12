package server

import (
    "github.com/Madduxv/mini-redis/internal/storage"
)

type Server struct {
    storage *storage.Storage
}

// NewServer initializes and returns a new Server instance.
func NewServer() *Server {
    return &Server{
        storage: storage.NewStorage(),
    }
}

// HandleHSet handles the HSET command from a client.
func (s *Server) HandleHSet(key, field, value string) {
    s.storage.HSet(key, field, value)
}

func (s *Server) HandleHGet(key, field string) (string, bool) {
    return s.storage.HGet(key, field)
}
