package in_memory_test

import (
	"sync"
	"testing"

	"github.com/bgoldovsky/service-rus-id/internal/domain/mock"
	"github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
	db "github.com/bgoldovsky/service-rus-id/internal/storage"
	. "github.com/bgoldovsky/service-rus-id/internal/storage/in_memory"
)

func TestNewInMemoryStorage_Success(t *testing.T) {
	store := make(map[valuetypes.UserID]*user.Snapshot)
	ma := &sync.RWMutex{}

	storage, err := NewInMemoryStorage(store, ma)

	if storage == nil {
		t.Errorf("expected: storage, act: %v", nil)
	}

	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}
}

func TestNewInMemoryStorage_Err(t *testing.T) {
	store := make(map[valuetypes.UserID]*user.Snapshot)
	ma := &sync.RWMutex{}

	storage, err := NewInMemoryStorage(nil, ma)

	if storage != nil {
		t.Errorf("expected: nil, act: %v", storage)
	}

	if err != ErrInvalidStore {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	storage, err = NewInMemoryStorage(store, nil)

	if storage != nil {
		t.Errorf("expected: nil, act: %v", storage)
	}

	if err != ErrInvalidMutex {
		t.Errorf("expected: %v, act: %v", nil, err)
	}
}

func TestInMemoryStorage_Get_Success(t *testing.T) {
	store := map[valuetypes.UserID]*user.Snapshot{
		*mock.UserID:      &mock.UserSnapshot,
		*mock.UserOtherID: &mock.UserOtherSnapshot,
	}
	ma := &sync.RWMutex{}

	storage, _ := NewInMemoryStorage(store, ma)

	snapshot, err := storage.Get(*mock.UserID)

	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	if snapshot.UserID != mock.UserIDRaw {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	snapshot, err = storage.Get(*mock.UserOtherID)

	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	if snapshot.UserID != mock.UserOtherIDRaw {
		t.Errorf("expected: %v, act: %v", nil, err)
	}
}

func TestInMemoryStorage_Get_NotFound(t *testing.T) {
	store := map[valuetypes.UserID]*user.Snapshot{
		*mock.UserID:      &mock.UserSnapshot,
		*mock.UserOtherID: &mock.UserOtherSnapshot,
	}
	ma := &sync.RWMutex{}

	storage, _ := NewInMemoryStorage(store, ma)

	snapshot, err := storage.Get(*valuetypes.CreateUserID())

	if err != db.ErrNotFound {
		t.Errorf("expected: %v, act: %v", db.ErrNotFound, err)
	}

	if snapshot != nil {
		t.Errorf("expected: %v, act: %v", nil, snapshot)
	}
}

func TestInMemoryStore_GetList_Success(t *testing.T) {
	store := map[valuetypes.UserID]*user.Snapshot{
		*mock.UserID:      &mock.UserSnapshot,
		*mock.UserOtherID: &mock.UserOtherSnapshot,
	}
	ma := &sync.RWMutex{}

	storage, _ := NewInMemoryStorage(store, ma)

	snapshots, err := storage.GetList(nil, nil)

	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	if act := len(snapshots); act != 2 {
		t.Errorf("expected: %v, act: %v", 2, act)
	}
}

func TestInMemoryStore_GetList_NotFound(t *testing.T) {
	store := map[valuetypes.UserID]*user.Snapshot{}
	ma := &sync.RWMutex{}

	storage, _ := NewInMemoryStorage(store, ma)

	snapshots, err := storage.GetList(nil, nil)

	if err != db.ErrNotFound {
		t.Errorf("expected: %v, act: %v", db.ErrNotFound, err)
	}

	if snapshots != nil {
		t.Errorf("expected: %v, act: %v", nil, snapshots)
	}
}
