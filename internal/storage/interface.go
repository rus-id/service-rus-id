package storage

import (
	"context"
	"errors"

	"github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

var (
	ErrNotFound = errors.New("user not found")
)

type UserRepository interface {
	Get(ctx context.Context, id valuetypes.UserID) (*user.Snapshot, error)
	GetList(ctx context.Context, offset, limit *int) ([]*user.Snapshot, error)
}
