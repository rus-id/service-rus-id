package in_memory_test

import (
	"reflect"
	"testing"

	"github.com/bgoldovsky/service-rus-id/internal/repository"

	"github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
	. "github.com/bgoldovsky/service-rus-id/internal/repository/in_memory"
	"github.com/bgoldovsky/service-rus-id/internal/repository/mock"
)

func TestNewInMemoryRepository_Success(t *testing.T) {
	store := make(map[valuetypes.UserID]*user.Snapshot)

	repo, err := NewInMemoryRepository(store)

	if repo == nil {
		t.Errorf("expected: repo, act: %v", nil)
	}

	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, repo)
	}
}

func TestNewInMemoryRepository_Err(t *testing.T) {
	repo, err := NewInMemoryRepository(nil)

	if repo != nil {
		t.Errorf("expected: %v, act: %v", nil, repo)
	}

	if err != ErrInvalidStore {
		t.Errorf("expected: %v, act: %v", ErrInvalidStore, repo)
	}
}

func TestInMemoryRepository_Save_Success(t *testing.T) {
	store := make(map[valuetypes.UserID]*user.Snapshot)

	repo, _ := NewInMemoryRepository(store)

	userAggregate := mock.UserAggregate

	err := repo.Save(&userAggregate)
	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}
}

func TestInMemoryRepository_Save_Nil(t *testing.T) {
	store := make(map[valuetypes.UserID]*user.Snapshot)

	repo, _ := NewInMemoryRepository(store)

	err := repo.Save(nil)
	if err != ErrInvalidAggregate {
		t.Errorf("expected: %v, act: %v", ErrInvalidAggregate, err)
	}
}

func TestInMemoryRepository_Save_NilAggregate(t *testing.T) {
	store := make(map[valuetypes.UserID]*user.Snapshot)

	repo, _ := NewInMemoryRepository(store)

	userAggregate := mock.UserNilAggregate

	err := repo.Save(&userAggregate)
	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}
}

func TestInMemoryRepository_IsExist_Success(t *testing.T) {
	store := make(map[valuetypes.UserID]*user.Snapshot)

	repo, _ := NewInMemoryRepository(store)

	userAggregate := mock.UserAggregate

	err := repo.Save(&userAggregate)
	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	ok, err := repo.IsExist(userAggregate.GetID())
	if !ok {
		t.Errorf("expected: %v, act: %v", true, ok)
	}

	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}
}

func TestInMemoryRepository_IsExist_Removed(t *testing.T) {
	store := make(map[valuetypes.UserID]*user.Snapshot)

	repo, _ := NewInMemoryRepository(store)

	userAggregate := mock.UserAggregate
	userAggregate.Remove()

	err := repo.Save(&userAggregate)
	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	ok, err := repo.IsExist(userAggregate.GetID())
	if ok {
		t.Errorf("expected: %v, act: %v", false, ok)
	}

	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}
}

func TestInMemoryRepository_IsExist_NotExist(t *testing.T) {
	store := make(map[valuetypes.UserID]*user.Snapshot)

	repo, _ := NewInMemoryRepository(store)

	userAggregate := mock.UserAggregate
	ok, err := repo.IsExist(userAggregate.GetID())
	if ok {
		t.Errorf("expected: %v, act: %v", false, ok)
	}

	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}
}

func TestInMemoryRepository_Find_Success(t *testing.T) {
	store := make(map[valuetypes.UserID]*user.Snapshot)

	repo, _ := NewInMemoryRepository(store)

	userAggregate := mock.UserAggregate

	err := repo.Save(&userAggregate)
	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	restored, err := repo.Find(userAggregate.GetID())
	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	if !reflect.DeepEqual(restored, &userAggregate) {
		t.Errorf("expected: %v, act: %v", restored, userAggregate)
	}
}

func TestInMemoryRepository_Find_Removed(t *testing.T) {
	store := make(map[valuetypes.UserID]*user.Snapshot)

	repo, _ := NewInMemoryRepository(store)

	userAggregate := mock.UserAggregate
	userAggregate.Remove()

	err := repo.Save(&userAggregate)
	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	restored, err := repo.Find(userAggregate.GetID())
	if err != repository.ErrNotFound {
		t.Errorf("expected: %v, act: %v", repository.ErrNotFound, err)
	}

	if restored != nil {
		t.Errorf("expected: %v, act: %v", nil, restored)
	}
}

func TestInMemoryRepository_Find_NotExist(t *testing.T) {
	store := make(map[valuetypes.UserID]*user.Snapshot)

	repo, _ := NewInMemoryRepository(store)

	userAggregate := mock.UserAggregate
	restored, err := repo.Find(userAggregate.GetID())
	if err != repository.ErrNotFound {
		t.Errorf("expected: %v, act: %v", repository.ErrNotFound, err)
	}

	if restored != nil {
		t.Errorf("expected: %v, act: %v", nil, restored)
	}
}
