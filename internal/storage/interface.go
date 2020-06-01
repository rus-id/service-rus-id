package storage

import (
	"errors"

	"github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

var (
	ErrNotFound = errors.New("user not found")
)

type UserRepository interface {
	Get(id valuetypes.UserID) (*user.Snapshot, error)
	GetList(start, limit *int) ([]*user.Snapshot, error)
}