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

func (s *Server) HandleHAdd(key, field, value string) {
	s.storage.HAdd(key, field, value)
}

func (s *Server) HandleHRem(key, field, value string) {
	s.storage.HRem(key, field, value)
}

func (s *Server) HandleRPush(key, field, value string) {
	s.storage.RPush(key, field, value)
}

func (s *Server) HandleLRange(key, field string, start, end int) []string {
	return s.storage.LRange(key, field, start, end)
}

// Deprecated
func (s *Server) HandleHSetList(key string, field string, value []string) {
	s.storage.HSetList(key, field, value)
}

// Deprecated
func (s *Server) HandleHGetList(key, field string) ([]string, bool) {
	return s.storage.HGetList(key, field)
}

func (s *Server) HandleHRemove(key string) {
	s.storage.HRemove(key)
}

func (s *Server) HandleSAdd(key, value string) {
	s.storage.SAdd(key, value)
}

func (s *Server) HandleSRem(key, value string) int8 {
	return s.storage.SRem(key, value)
}

func (s *Server) HandleSGet(key string) ([]string, bool) {
	return s.storage.SGet(key)
}

// Deprecated
func (s *Server) HandleHRemoveListField(key string, field string) bool {
	return s.storage.HRemoveListField(key, field)
}

func (s *Server) HandleHRemoveStringField(key string, field string) bool {
	return s.storage.HRemoveStringField(key, field)
}
