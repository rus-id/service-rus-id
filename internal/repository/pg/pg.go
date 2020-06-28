package pg

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v4"

	"github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

type UserRepositoryPg struct {
	database Queryer
}

type Queryer interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

func NewInMemoryRepository(database Queryer) *UserRepositoryPg {
	return &UserRepositoryPg{
		database: database,
	}
}

func (r *UserRepositoryPg) ForUpdate(ctx context.Context, id valuetypes.UserID) (*user.User, error) {
	query := `select "id", "first_name", "middle_name", last_name from "users" where "id" = $1 for update`

	var userID uuid.UUID
	var firstName, lastName string
	var middleName *string

	err := r.database.QueryRow(ctx, query, id).Scan(&userID, &firstName, middleName, &lastName)

	snapshot := user.NewSnapshot(
		userID,
		firstName,
		middleName,
		lastName,
		0,
		"",
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		0,
		0,
		0,
		nil,
		0,
		false,
		0,
		0)

	aggregate, err := user.LoadFromSnapshot(&snapshot)
	if err != nil {
		return nil, err
	}

	return aggregate, err
}

// TODO: реализовать методы репозитория

func (r *UserRepositoryPg) Find(id valuetypes.UserID) (*user.User, error) {
	return nil, errors.New("not implemented")
}

func (r *UserRepositoryPg) Save(user *user.User) error {
	return errors.New("not implemented")
}
func (r *UserRepositoryPg) IsExist(id valuetypes.UserID) (bool, error) {
	return false, errors.New("not implemented")
}
