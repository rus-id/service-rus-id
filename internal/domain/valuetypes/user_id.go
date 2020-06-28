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

func NewUserID(value uuid.UUID) (*UserID, error) {
	userID := UserID(value)
	return &userID, nil
}

func (u UserID) String() string {
	return uuid.UUID(u).String()
}
