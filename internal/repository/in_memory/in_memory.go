package in_memory

import (
	"errors"
	"sync"
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/repository"
)

var (
	ErrInvalidStore     = errors.New("in memory store not specified")
	ErrInvalidAggregate = errors.New("invalid aggregate")
)

type InMemoryRepository struct {
	store map[valuetypes.UserID]*user.Snapshot
	ma    sync.RWMutex
}

func NewInMemoryRepository(store map[valuetypes.UserID]*user.Snapshot) (*InMemoryRepository, error) {
	if store == nil {
		return nil, ErrInvalidStore
	}

	return &InMemoryRepository{
		store: store,
	}, nil
}

func (r *InMemoryRepository) Find(id valuetypes.UserID) (*user.User, error) {
	r.ma.RLock()
	defer r.ma.RUnlock()

	val, ok := r.store[id]
	if !ok || val.IsRemoved {
		return nil, repository.ErrNotFound
	}

	return user.LoadFromSnapshot(val)
}

func (r *InMemoryRepository) Save(u user.Aggregate) error {
	if u == nil {
		return ErrInvalidAggregate
	}

	if _, ok := u.(*user.Nil); ok {
		return nil
	}

	userAggregate := u.(*user.User)

	r.ma.Lock()
	defer r.ma.Unlock()

	snapshot, err := user.GetSnapshot(userAggregate, time.Now())
	if err != nil {
		return err
	}
	r.store[u.GetID()] = snapshot

	return nil
}
func (r *InMemoryRepository) IsExist(id valuetypes.UserID) (bool, error) {
	r.ma.RLock()
	defer r.ma.RUnlock()

	val, ok := r.store[id]

	return ok && !val.IsRemoved, nil
}
