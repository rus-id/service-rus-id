package repository

import (
	"errors"

	"github.com/bgoldovsky/service-rus-id/internal/domain/aggregates"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

type UserRepositoryPg struct {
}

func (r *UserRepositoryPg) FindOrCreate(id valuetypes.UserID) (*aggregates.User, error) {
	return nil, errors.New("not implemented")
}

func (r *UserRepositoryPg) Save(user *aggregates.User) error {
	return errors.New("not implemented")
}
func (r *UserRepositoryPg) IsExist(id valuetypes.UserID) (bool, error) {
	return false, errors.New("not implemented")
}
