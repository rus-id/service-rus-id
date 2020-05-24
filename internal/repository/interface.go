package repository

import (
	"github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

type UserRepository interface {
	FindOrCreate(id valuetypes.UserID) (*user.User, error)
	Save(user *user.User) error
	IsExist(id valuetypes.UserID) (bool, error)
}
