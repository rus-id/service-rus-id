package aggregates

import (
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"

	"github.com/bgoldovsky/service-rus-id/internal/domain/snapshots"

	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

type UserAggregate interface {
	GetSnapshot() *snapshots.UserSnapshot
}

type User struct {
	ID               valuetypes.UserID
	Phone            *valuetypes.Phone
	Passport         *passport.Passport
	DrivingLicense   *driving_license.DrivingLicense
	Snils            valuetypes.Snils
	Inn              valuetypes.Inn
	Photo            valuetypes.Photo
	Card             valuetypes.Card
	RegistrationDate time.Time
	Rating           *valuetypes.Rating
	Status           valuetypes.UserState
	IsRemoved        bool
	Version          int64
}

func NewUser(
	ID valuetypes.UserID,
	snils valuetypes.Snils) *User {
	return &User{
		ID:    ID,
		Snils: snils,
	}
}

func NewUserFromSnapshot(snapshot snapshots.UserSnapshot) *User {
	return nil
}
