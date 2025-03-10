package session

import (
	"sync"
	"time"
)

type Data struct {
	Data      map[string]interface{}
	ExpiresAt time.Time
}

type Store struct {
	sessions map[string]*Data
	mu       sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		sessions: make(map[string]*Data),
	}
}

func (s *Store) Set(sessionID, key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if session, ok := s.sessions[sessionID]; ok {
		session.Data[key] = value
	} else {
		s.sessions[sessionID] = &Data{
			Data:      map[string]interface{}{key: value},
			ExpiresAt: time.Now().Add(30 * time.Minute),
		}
	}
}

func (s *Store) Get(sessionID, key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if session, ok := s.sessions[sessionID]; ok {
		return session.Data[key], true
	}
	return nil, false
}

func (s *Store) Delete(sessionID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, sessionID)
}
