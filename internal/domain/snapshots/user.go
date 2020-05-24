package snapshots

import "github.com/google/uuid"

type UserSnapshot struct {
	UserID           uuid.UUID
	FirstName        string
	MiddleName       *string
	LastName         string
	CountryCode      int64
	PhoneNumber      string
	Passport         *PassportSnapshot
	DrivingLicense   *DrivingLicenseSnapshot
	Snils            *string
	Inn              *string
	Photo            []byte
	CardNumber       *string
	CardExpired      *int64
	RegistrationDate int64
	RatingPositive   int64
	RatingNegative   int64
	Tolerances       []ToleranceSnapshot
	State            int64
	IsRemoved        bool
	Version          int64
}

func NewUser(
	userID uuid.UUID,
	firstName string,
	middleName *string,
	lastName string,
	countryCode int64,
	phoneNumber string,
	passport *PassportSnapshot,
	drivingLicense *DrivingLicenseSnapshot,
	snils *string,
	inn *string,
	photo []byte,
	cardNumber *string,
	cardExpired *int64,
	registrationDate int64,
	ratingPositive int64,
	ratingNegative int64,
	tolerances []ToleranceSnapshot,
	state int64,
	isRemoved bool,
	version int64,
) UserSnapshot {
	return UserSnapshot{
		UserID:           userID,
		FirstName:        firstName,
		MiddleName:       middleName,
		LastName:         lastName,
		CountryCode:      countryCode,
		PhoneNumber:      phoneNumber,
		Passport:         passport,
		DrivingLicense:   drivingLicense,
		Snils:            snils,
		Inn:              inn,
		Photo:            photo,
		CardNumber:       cardNumber,
		CardExpired:      cardExpired,
		RegistrationDate: registrationDate,
		RatingPositive:   ratingPositive,
		RatingNegative:   ratingNegative,
		Tolerances:       tolerances,
		State:            state,
		IsRemoved:        isRemoved,
		Version:          version,
	}
}
