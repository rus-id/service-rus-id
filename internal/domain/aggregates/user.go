package aggregates

import (
	"errors"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license"
	dlValueTypes "github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"
	passValueTypes "github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/domain/snapshots"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
	"github.com/google/uuid"
	"reflect"
	"time"
)

var (
	ErrInvalidID               = errors.New("user aggregate invalid ID")
	ErrInvalidName             = errors.New("user aggregate invalid name")
	ErrInvalidPhone            = errors.New("user aggregate invalid phone")
	ErrInvalidRating           = errors.New("user aggregate invalid rating")
	ErrInvalidState            = errors.New("user aggregate invalid state")
	ErrInvalidRegistrationDate = errors.New("user aggregate invalid registration date")
)

type UserAggregate interface {
	GetID() valuetypes.UserID
	IsRemoved() bool
	GetSnapshot(timestamp time.Time) snapshots.UserSnapshot

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

	if rating != nil {
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
	}, nil
}

func LoadUser(snapshot snapshots.UserSnapshot) (*User, error) {
	id := valuetypes.UserID(snapshot.UserID)
	name, err := valuetypes.NewName(snapshot.FirstName, snapshot.MiddleName, snapshot.LastName)
	if err != nil {
		return nil, err
	}

	phone, err := valuetypes.NewPhone(valuetypes.CountryCode(snapshot.CountryCode), snapshot.PhoneNumber)
	if err != nil {
		return nil, err
	}

	rating, err := valuetypes.NewRating(int(snapshot.RatingPositive), int(snapshot.RatingNegative))
	if err != nil {
		return nil, err
	}

	registrationDate := time.Unix(snapshot.RegistrationDate, 0)
	state := valuetypes.UserState(snapshot.State)

	user, err := NewUser(&id, name, phone, &registrationDate, rating, state)
	if err != nil {
		return nil, err
	}

	var pass *passport.Passport
	if snapshot.Passport != nil {
		id, err := passValueTypes.NewPassportID(snapshot.Passport.Serial, snapshot.Passport.Number)
		if err != nil {
			return nil, err
		}

		name, err := valuetypes.NewName(snapshot.Passport.FirstName, snapshot.Passport.MiddleName, snapshot.Passport.LastName)
		if err != nil {
			return nil, err
		}

		birthday := time.Unix(snapshot.Passport.Birthday, 0)
		issueDate := time.Unix(snapshot.Passport.IssuedDate, 0)

		issue, err := passValueTypes.NewPassportIssue(snapshot.Passport.IssuedOrganisation, issueDate, snapshot.Passport.IssuedCode)
		if err != nil {
			return nil, err
		}

		registration, err := valuetypes.NewAddress(snapshot.Passport.Registration)
		if err != nil {
			return nil, err
		}

		validation := passValueTypes.NewPassportValidation(
			snapshot.Passport.UfmsValidation,
			snapshot.Passport.MvdValidation,
			snapshot.Passport.FsspValidation,
			snapshot.Passport.DocumentValidation)

		user.passport, err = passport.NewPassport(id, name, &birthday, issue, registration, validation)
		if err != nil {
			return nil, err
		}
	}

	var drivingLicense *driving_license.DrivingLicense
	if snapshot.DrivingLicense != nil {
		id, err := dlValueTypes.NewDrivingLicenseID(snapshot.DrivingLicense.Serial, snapshot.DrivingLicense.Number)
		if err != nil {
			return nil, err
		}

		name, err := valuetypes.NewName(snapshot.DrivingLicense.FirstName, snapshot.DrivingLicense.MiddleName, snapshot.DrivingLicense.LastName)
		if err != nil {
			return nil, err
		}

		category := dlValueTypes.DrivingLicenseCategory(snapshot.DrivingLicense.Category)
		birthday := time.Unix(snapshot.DrivingLicense.Birthday, 0)
		issued := time.Unix(snapshot.DrivingLicense.Issued, 0)
		expired := time.Unix(snapshot.DrivingLicense.Expired, 0)

		user.drivingLicense, err := driving_license.NewDrivingLicense(id, category, name, &birthday, &issued, &expired,)
		if err != nil {
			return nil, err
		}
	}
	user.drivingLicense = drivingLicense

}

func (u *User) GetID() valuetypes.UserID {
	return u.id
}

func (u *User) IsRemoved() bool {
	return u.isRemoved
}

func (u *User) GetSnapshot(timestamp time.Time) snapshots.UserSnapshot {
	var passport *snapshots.PassportSnapshot
	if u.passport != nil {
		tmp := snapshots.NewPassport(
			u.passport.GetID().GetSerial(),
			u.passport.GetID().GetNumber(),
			u.passport.GetName().GetFirst(),
			u.passport.GetName().GetMiddle(),
			u.passport.GetName().GetLast(),
			u.passport.GetBirthday().Unix(),
			u.passport.GetIssued().GetOrganisation(),
			u.passport.GetIssued().GetDate().Unix(),
			u.passport.GetIssued().GetCode(),
			string(u.passport.GetRegistration()),
			u.passport.GetValidation().GetUfms(),
			u.passport.GetValidation().GetMvd(),
			u.passport.GetValidation().GetFssp(),
			u.passport.GetValidation().GetDocument(),
			timestamp.Unix(),
		)
		passport = &tmp
	}

	var drivingLicense *snapshots.DrivingLicenseSnapshot
	if u.drivingLicense != nil {
		tmp := snapshots.NewDrivingLicense(
			u.drivingLicense.GetID().GetSerial(),
			u.drivingLicense.GetID().GetNumber(),
			int64(u.drivingLicense.GetCategory()),
			u.drivingLicense.GetName().GetFirst(),
			u.drivingLicense.GetName().GetMiddle(),
			u.drivingLicense.GetName().GetLast(),
			u.drivingLicense.GetBirthday().Unix(),
			u.drivingLicense.GetIssued().Unix(),
			u.drivingLicense.GetExpired().Unix(),
			u.drivingLicense.GetResidence().GetValue(),
			u.drivingLicense.GetSpecialMarks(),
			u.drivingLicense.GetValidation().GetGibdd(),
			u.drivingLicense.GetValidation().GetDocument(),
			timestamp.Unix(),
		)
		drivingLicense = &tmp
	}

	var snils *string
	if u.snils != nil {
		tmp := string(*u.snils)
		snils = &tmp
	}

	var inn *string
	if u.inn != nil {
		tmp := string(*u.inn)
		inn = &tmp
	}

	var cardNumber *string
	var cardExpired *int64
	if u.card != nil {
		tmpNum := u.card.GetNumber()
		tmpExp := u.card.GetExpired().Unix()
		cardNumber = &tmpNum
		cardExpired = &tmpExp
	}

	return snapshots.NewUser(
		uuid.UUID(u.id),
		u.name.GetFirst(),
		u.name.GetMiddle(),
		u.name.GetLast(),
		int64(u.phone.GetCode()),
		u.phone.GetNumber(),
		passport,
		drivingLicense,
		snils,
		inn,
		u.photo,
		cardNumber,
		cardExpired,
		u.registrationDate.Unix(),
		int64(u.rating.GetPositive()),
		int64(u.rating.GetNegative()),
		[]snapshots.ToleranceSnapshot{},
		int64(u.state),
		u.isRemoved,
		u.version)
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
