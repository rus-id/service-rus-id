package aggregates

import (
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"

	"github.com/bgoldovsky/service-rus-id/internal/domain/snapshots"

	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

type UserAggregate interface {
	GetID() valuetypes.UserID
	GetSnapshot() *snapshots.UserSnapshot
}

type User struct {
	id               *valuetypes.UserID
	phone            *valuetypes.Phone
	passport         *passport.Passport
	drivingLicense   *driving_license.DrivingLicense
	snils            *valuetypes.Snils
	inn              *valuetypes.Inn
	photo            *valuetypes.Photo
	card             *valuetypes.Card
	registrationDate *time.Time
	rating           *valuetypes.Rating
	tolerances       map[valuetypes.UserID]*valuetypes.Tolerance
	status           valuetypes.UserState
	isRemoved        bool
	version          int64
}

func NewUser(
	id *valuetypes.UserID,
	snils *valuetypes.Snils) *User {
	return &User{
		id:    id,
		snils: snils,
	}
}

func NewUserFromSnapshot(snapshot snapshots.UserSnapshot) *User {
	return nil
}
