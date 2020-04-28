package repository

import (
	"github.com/bgoldovsky/service-rus-id/internal/domain/aggregates"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

type UserRepository interface {
	FindOrCreate(id valuetypes.UserID) (*aggregates.User, error)
	Save(user *aggregates.User) error
	IsExist(id valuetypes.UserID) bool
}
