package valuetypes

import "github.com/google/uuid"

type UserID uuid.UUID

func CreateUserID() UserID {
	ID := uuid.New()
	return UserID(ID)
}

func NewUserID(value string) (UserID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return UserID{}, err
	}

	return UserID(id), nil
}
