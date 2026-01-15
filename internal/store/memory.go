package store

import (
	"sync"
	"time"

	"github.com/HallenWeaver/go-grpc-mini/internal/service"
	"github.com/google/uuid"
)

type MemoryStore struct {
	mu    sync.Mutex
	users map[string]service.User
}

func New() *MemoryStore {
	return &MemoryStore{
		users: make(map[string]service.User),
	}
}

func (s *MemoryStore) Create(user service.User) (service.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user.ID = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Active = true

	s.users[user.ID] = user
	return user, nil
}

func (s *MemoryStore) GetByID(id string) (service.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return service.User{}, service.ErrNotFound
	}
	return user, nil
}

func (s *MemoryStore) GetByEmail(email string) (service.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.users {
		if user.Email == email {
			return user, nil
		}
	}
	return service.User{}, service.ErrNotFound
}

func (s *MemoryStore) Update(id string, displayName *string) (service.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return service.User{}, service.ErrNotFound
	}

	if displayName != nil {
		user.DisplayName = *displayName
	}

	user.UpdatedAt = time.Now()

	s.users[id] = user
	return user, nil
}
