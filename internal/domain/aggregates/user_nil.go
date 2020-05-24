package aggregates

import (
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"
	"github.com/bgoldovsky/service-rus-id/internal/domain/snapshots"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
	"github.com/google/uuid"
)

type UserNil struct {
	id valuetypes.UserID
}

func NewUserNil(id *valuetypes.UserID) (*UserNil, error) {
	if id == nil {
		return nil, ErrInvalidID
	}

	return &UserNil{id: *id}, nil
}

func (u *UserNil) GetID() valuetypes.UserID {
	return u.id
}

func (u *UserNil) IsRemoved() bool {
	return true
}

func (u *UserNil) GetSnapshot() *snapshots.UserSnapshot {
	return &snapshots.UserSnapshot{
		UserID:    uuid.UUID(u.id),
		IsRemoved: true,
	}
}

func (u *UserNil) ChangeName(_ *valuetypes.Name) {
}

func (u *UserNil) ChangePhone(_ *valuetypes.Phone) {
}

func (u *UserNil) ChangePassport(_ *passport.Passport) {
}

func (u *UserNil) ChangeDrivingLicense(_ *driving_license.DrivingLicense) {
}

func (u *UserNil) ChangeSnils(_ *valuetypes.Snils) {
}

func (u *UserNil) ChangeInn(_ *valuetypes.Inn) {
}

func (u *UserNil) ChangePhoto(_ *valuetypes.Photo) {
}

func (u *UserNil) ChangeCard(_ *valuetypes.Card) {
}

func (u *UserNil) IncreaseRating() {
}

func (u *UserNil) DecreaseRating() {
}

func (u *UserNil) GrantAccess(_ valuetypes.UserID, _ valuetypes.Accessor) {
}

func (u *UserNil) RevokeAccess(_ valuetypes.UserID, _ valuetypes.Accessor) {
}

func (u *UserNil) GrantFullAccess(_ valuetypes.UserID) {
}

func (u *UserNil) RevokeFullAccess(_ valuetypes.UserID) {
}

func (u *UserNil) Block() {
}

func (u *UserNil) Activate() {
}

func (u *UserNil) Remove() {
}
