package pg

import (
	"errors"

	"github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

type UserRepositoryPg struct {
}

func (r *UserRepositoryPg) Find(id valuetypes.UserID) (*user.User, error) {
	return nil, errors.New("not implemented")
}

func (r *UserRepositoryPg) Save(user *user.User) error {
	return errors.New("not implemented")
}
func (r *UserRepositoryPg) IsExist(id valuetypes.UserID) (bool, error) {
	return false, errors.New("not implemented")
}
