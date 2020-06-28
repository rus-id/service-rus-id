package repository

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
	Find(ctx context.Context, id valuetypes.UserID) (*user.User, error)
	Save(ctx context.Context, user *user.User) error
	IsExist(ctx context.Context, id valuetypes.UserID) (bool, error)
}
