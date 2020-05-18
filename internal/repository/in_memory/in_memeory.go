package in_memory

import (
	"errors"

	"github.com/bgoldovsky/service-rus-id/internal/domain/aggregates"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

type UserRepositoryInMemory struct {
}

func (r *UserRepositoryInMemory) FindOrCreate(id valuetypes.UserID) (*aggregates.User, error) {
	return nil, errors.New("not implemented")
}

func (r *UserRepositoryInMemory) Save(user *aggregates.User) error {
	return errors.New("not implemented")
}
func (r *UserRepositoryInMemory) IsExist(id valuetypes.UserID) (bool, error) {
	return false, errors.New("not implemented")
}
