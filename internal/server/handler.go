package server

import (
  "github.com/Madduxv/mini-redis/internal/storage"
)

type Server struct {
  storage *storage.Storage
}

func NewServer() *Server {
  return &Server{
    storage: storage.NewStorage(),
  }
}

func (s *Server) HandleHSet(key, field, value string) {
  s.storage.HSet(key, field, value)
}

func (s *Server) HandleHGet(key, field string) (string, bool) {
  return s.storage.HGet(key, field)
}

func (s *Server) HandleRPush(key, field, value string) {
  s.storage.RPush(key, field, value)
}

func (s *Server) HandleHSetList(key string, field string, value []string) {
  s.storage.HSetList(key, field, value)
}

func (s *Server) HandleHGetList(key, field string) ([]string, bool) {
  return s.storage.HGetList(key, field)
}

func (s *Server) HandleHRemove(key string) {
  s.storage.HRemove(key)
}

func (s *Server) HandleHRemoveListField(key string, field string) bool {
  return s.storage.HRemoveListField(key, field)
}

func (s *Server) HandleHRemoveStringField(key string, field string) bool {
  return s.storage.HRemoveStringField(key, field)
}
