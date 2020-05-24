package user

import (
	"errors"
	"reflect"
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

var (
	ErrInvalidID               = errors.New("user aggregate invalid ID")
	ErrInvalidName             = errors.New("user aggregate invalid name")
	ErrInvalidPhone            = errors.New("user aggregate invalid phone")
	ErrInvalidRating           = errors.New("user aggregate invalid rating")
	ErrInvalidState            = errors.New("user aggregate invalid state")
	ErrInvalidRegistrationDate = errors.New("user aggregate invalid registration date")
)

type Aggregate interface {
	GetID() valuetypes.UserID
	IsRemoved() bool

	ChangeName(name *valuetypes.Name)
	ChangePhone(phone *valuetypes.Phone)
	ChangePassport(pass *passport.Passport)
	ChangeDrivingLicense(drivingLicense *driving_license.DrivingLicense)
	ChangeSnils(snils *valuetypes.Snils)
	ChangeInn(inn *valuetypes.Inn)
	ChangePhoto(photo *valuetypes.Photo)
	ChangeCard(card *valuetypes.Card)

	IncreaseRating()
	DecreaseRating()

	GrantAccess(userID valuetypes.UserID, accessor valuetypes.Accessor)
	RevokeAccess(userID valuetypes.UserID, accessor valuetypes.Accessor)
	GrantFullAccess(userID valuetypes.UserID)
	RevokeFullAccess(userID valuetypes.UserID)

	Block()
	Activate()
	Remove()
}

type User struct {
	id               valuetypes.UserID
	name             valuetypes.Name
	phone            valuetypes.Phone
	passport         *passport.Passport
	drivingLicense   *driving_license.DrivingLicense
	snils            *valuetypes.Snils
	inn              *valuetypes.Inn
	photo            valuetypes.Photo
	card             *valuetypes.Card
	registrationDate time.Time
	rating           valuetypes.Rating
	tolerances       map[valuetypes.UserID]valuetypes.Tolerance
	state            valuetypes.UserState
	isRemoved        bool
	version          int64
}

func NewUser(
	id *valuetypes.UserID,
	name *valuetypes.Name,
	phone *valuetypes.Phone,
	registrationDate *time.Time,
	rating *valuetypes.Rating,
	state valuetypes.UserState,
	isRemoved bool,
	version int64,
) (*User, error) {
	if id == nil {
		return nil, ErrInvalidID
	}

	if name == nil {
		return nil, ErrInvalidName
	}

	if phone == nil {
		return nil, ErrInvalidPhone
	}

	if registrationDate == nil {
		return nil, ErrInvalidRegistrationDate
	}

	if rating == nil {
		return nil, ErrInvalidRating
	}

	if !state.IsValid() {
		return nil, ErrInvalidState
	}

	return &User{
		id:               *id,
		name:             *name,
		phone:            *phone,
		registrationDate: *registrationDate,
		rating:           *rating,
		tolerances:       make(map[valuetypes.UserID]valuetypes.Tolerance),
		state:            state,
		isRemoved:        isRemoved,
		version:          version,
	}, nil
}

func (u *User) GetID() valuetypes.UserID {
	return u.id
}

func (u *User) IsRemoved() bool {
	return u.isRemoved
}

func (u *User) ChangeName(name *valuetypes.Name) {
	if name == nil {
		return
	}

	if *name == u.name {
		return
	}

	u.name = *name
}

func (u *User) ChangePhone(phone *valuetypes.Phone) {
	if phone == nil {
		return
	}

	if *phone == u.phone {
		return
	}

	u.phone = *phone
}

func (u *User) ChangePassport(pass *passport.Passport) {
	if pass == nil {
		return
	}

	if u.passport == nil {
		u.passport = pass
		return
	}

	if *u.passport == *pass {
		return
	}

	u.passport = pass
}

func (u *User) ChangeDrivingLicense(drivingLicense *driving_license.DrivingLicense) {
	if drivingLicense == nil {
		return
	}

	if u.drivingLicense == nil {
		u.drivingLicense = drivingLicense
		return
	}

	if *u.drivingLicense == *drivingLicense {
		return
	}

	u.drivingLicense = drivingLicense
}

func (u *User) ChangeSnils(snils *valuetypes.Snils) {
	if snils == nil {
		return
	}

	if u.snils == nil {
		u.snils = snils
		return
	}

	if *u.snils == *snils {
		return
	}

	u.snils = snils
}

func (u *User) ChangeInn(inn *valuetypes.Inn) {
	if inn == nil {
		return
	}

	if u.inn == nil {
		u.inn = inn
		return
	}

	if *u.inn == *inn {
		return
	}

	u.inn = inn
}

func (u *User) ChangePhoto(photo *valuetypes.Photo) {
	if photo == nil {
		return
	}

	if u.photo == nil {
		u.photo = *photo
		return
	}

	if reflect.DeepEqual(u.photo, photo) {
		return
	}

	u.photo = *photo
}

func (u *User) ChangeCard(card *valuetypes.Card) {
	if card == nil {
		return
	}

	if u.card == nil {
		u.card = card
		return
	}

	if *u.card == *card {
		return
	}

	u.card = card
}

func (u *User) IncreaseRating() {
	u.rating.AddPositive()
}

func (u *User) DecreaseRating() {
	u.rating.AddNegative()
}

func (u *User) GrantAccess(userID valuetypes.UserID, accessor valuetypes.Accessor) {
	tolerance, ok := u.tolerances[userID]
	if ok {
		tolerance = tolerance.AddAccess(accessor)
		u.tolerances[userID] = tolerance
		return
	}

	newTolerances, err := valuetypes.NewTolerance(&userID, []valuetypes.Accessor{accessor})
	if err != nil {
		return
	}

	u.tolerances[userID] = *newTolerances
}

func (u *User) RevokeAccess(userID valuetypes.UserID, accessor valuetypes.Accessor) {
	tolerance, ok := u.tolerances[userID]
	if !ok {
		return
	}

	tolerance = tolerance.RemoveAccess(accessor)
	u.tolerances[userID] = tolerance
}

func (u *User) GrantFullAccess(userID valuetypes.UserID) {
	tolerance, ok := u.tolerances[userID]
	if ok {
		tolerance = tolerance.AddFullAccess()
		u.tolerances[userID] = tolerance
		return
	}

	newTolerances, err := valuetypes.NewTolerance(&userID, []valuetypes.Accessor{})
	if err != nil {
		return
	}

	u.tolerances[userID] = newTolerances.AddFullAccess()
}

func (u *User) RevokeFullAccess(userID valuetypes.UserID) {
	_, ok := u.tolerances[userID]
	if !ok {
		return
	}

	delete(u.tolerances, userID)
}

func (u *User) Block() {
	if u.state == valuetypes.UserStateBlocked {
		return
	}

	u.state = valuetypes.UserStateBlocked
}

func (u *User) Activate() {
	if u.state == valuetypes.UserStateActive {
		return
	}

	u.state = valuetypes.UserStateActive
}

func (u *User) Remove() {
	if u.isRemoved {
		return
	}

	u.isRemoved = true
}
