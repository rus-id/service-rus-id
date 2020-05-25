package repository

import (
	"errors"

	"github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

var (
	ErrNotFound = errors.New("user not fount")
)

type UserRepository interface {
	Find(id valuetypes.UserID) (*user.User, error)
	Save(user *user.User) error
	IsExist(id valuetypes.UserID) (bool, error)
}
