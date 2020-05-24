package user

import (
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

type Nil struct {
	id valuetypes.UserID
}

func NewUserNil(id *valuetypes.UserID) (*Nil, error) {
	if id == nil {
		return nil, ErrInvalidID
	}

	return &Nil{id: *id}, nil
}

func (u *Nil) GetID() valuetypes.UserID {
	return u.id
}

func (u *Nil) IsRemoved() bool {
	return true
}

func (u *Nil) ChangeName(_ *valuetypes.Name) {
}

func (u *Nil) ChangePhone(_ *valuetypes.Phone) {
}

func (u *Nil) ChangePassport(_ *passport.Passport) {
}

func (u *Nil) ChangeDrivingLicense(_ *driving_license.DrivingLicense) {
}

func (u *Nil) ChangeSnils(_ *valuetypes.Snils) {
}

func (u *Nil) ChangeInn(_ *valuetypes.Inn) {
}

func (u *Nil) ChangePhoto(_ *valuetypes.Photo) {
}

func (u *Nil) ChangeCard(_ *valuetypes.Card) {
}

func (u *Nil) IncreaseRating() {
}

func (u *Nil) DecreaseRating() {
}

func (u *Nil) GrantAccess(_ valuetypes.UserID, _ valuetypes.Accessor) {
}

func (u *Nil) RevokeAccess(_ valuetypes.UserID, _ valuetypes.Accessor) {
}

func (u *Nil) GrantFullAccess(_ valuetypes.UserID) {
}

func (u *Nil) RevokeFullAccess(_ valuetypes.UserID) {
}

func (u *Nil) Block() {
}

func (u *Nil) Activate() {
}

func (u *Nil) Remove() {
}
