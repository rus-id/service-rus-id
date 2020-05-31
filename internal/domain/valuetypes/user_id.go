package valuetypes

import (
	"github.com/google/uuid"
)

type UserID uuid.UUID

func CreateUserID() *UserID {
	id := uuid.New()

	userID := UserID(id)
	return &userID
}

func NewUserID(value string) (*UserID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return nil, err
	}

	userID := UserID(id)
	return &userID, nil
}

func (u UserID) String() string {
	return uuid.UUID(u).String()
}
