package in_memory

import (
	"errors"
	"sync"

	"github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/storage"
)

var (
	ErrInvalidStore = errors.New("in memory storage not specified")
	ErrInvalidMutex = errors.New("mutex not specified")
)

type InMemoryStorage struct {
	store map[valuetypes.UserID]*user.Snapshot
	ma    sync.RWMutex
}

func NewInMemoryStorage(store map[valuetypes.UserID]*user.Snapshot, ma *sync.RWMutex) (*InMemoryStorage, error) {
	if store == nil {
		return nil, ErrInvalidStore
	}

	if ma == nil {
		return nil, ErrInvalidMutex
	}

	return &InMemoryStorage{
		store: store,
	}, nil
}

func (s *InMemoryStorage) Get(id valuetypes.UserID) (*user.Snapshot, error) {
	s.ma.RLock()
	defer s.ma.RUnlock()

	val, ok := s.store[id]
	if !ok || val.IsRemoved {
		return nil, storage.ErrNotFound
	}

	return val, nil
}

func (s *InMemoryStorage) GetList(_, _ *int) ([]*user.Snapshot, error) {
	s.ma.RLock()
	defer s.ma.RUnlock()

	if len(s.store) == 0 {
		return nil, storage.ErrNotFound
	}

	users := make([]*user.Snapshot, 0, len(s.store))
	for _, val := range s.store {
		users = append(users, val)
	}

	return users, nil
}
