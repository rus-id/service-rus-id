package mock

import (
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license"
	dlValues "github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"
	passportTypes "github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/domain/user"
	commonTypes "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

//Passport

const (
	passportSerial             = "7777"
	passportNumber             = "777777"
	passportFirstName          = "Boris"
	passportLastName           = "Goldovsky"
	passportUfmsValidation     = true
	passportMvdValidation      = false
	passportFsspValidation     = true
	passportAddress            = "Russia, Moscow"
	passportIssuedOrganisation = "MVD"
	passportIssuedCode         = "770-77"
	passportDocumentValidation = false
)

var (
	passportBirthday        = time.Date(1986, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	passportIssuedDate      = time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	passportID, _           = passportTypes.NewPassportID(passportSerial, passportNumber)
	passportName, _         = commonTypes.NewName(passportFirstName, nil, passportLastName)
	passportIssued, _       = passportTypes.NewPassportIssue(passportIssuedOrganisation, passportIssuedDate, passportIssuedCode)
	passportRegistration, _ = commonTypes.NewAddress(passportAddress)
	passportValidation      = passportTypes.NewPassportValidation(passportUfmsValidation, passportMvdValidation, passportFsspValidation, passportDocumentValidation)

	passportEntity, _ = passport.NewPassport(
		passportID,
		passportName,
		&passportBirthday,
		passportIssued,
		passportRegistration,
		passportValidation)
)

// Driving License

const (
	dlSerial             = "7777"
	dlNumber             = "777777"
	dlFirstName          = "Boris"
	dlLastName           = "Goldovsky"
	dlCountry            = "Russia"
	dlGibddValidation    = true
	dlDocumentValidation = false
	dlSpecialMarks       = "empty mark"
)

var (
	dlBirthday             = time.Date(1986, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	dlIssued               = time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	dlExpires              = time.Date(2025, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	dlMiddleName   *string = nil
	dlID, _                = dlValues.NewDrivingLicenseID(dlSerial, dlNumber)
	dlCategory             = dlValues.DrivingLicenseA
	dlName, _              = commonTypes.NewName(dlFirstName, dlMiddleName, dlLastName)
	dlResidence, _         = dlValues.NewResidence(dlCountry)
	dlValidation           = dlValues.NewDrivingLicenseValidation(dlGibddValidation, dlDocumentValidation)

	drivingLicenseEntity, _ = driving_license.NewDrivingLicense(
		dlID,
		dlCategory,
		dlName,
		&dlBirthday,
		&dlIssued,
		&dlExpires,
		dlResidence,
		dlSpecialMarks,
		dlValidation)
)

// User

const (
	userIDRaw          = "059b4e12-6983-4806-bd5a-cc3433e78f66"
	userOtherIDRaw     = "059b4e12-6983-4806-bd5a-cc3433e78f66"
	userCountryCode    = 7
	userPhoneNumber    = "9039615322"
	userNegativeRating = 40
	userPositiveRating = 50
	userStateRaw       = 1
	userIsRemoved      = false
	userVersion        = 12345
)

var (
	userCardExpires      = time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	userRegistrationDate = time.Date(2019, time.Month(3), 9, 1, 10, 30, 0, time.UTC)
	userInnRaw           = "926902267890"
	userSnilsRaw         = "59650418527"
	userCardNumber       = "4444333322221111"
	userPhotoRaw         = []byte{10, 20, 30, 40, 50}
	userID, _            = commonTypes.NewUserID(userIDRaw)
	userOtherID, _       = commonTypes.NewUserID(userOtherIDRaw)
	userName, _          = commonTypes.NewName(dlFirstName, dlMiddleName, dlLastName)
	userPhone, _         = commonTypes.NewPhone(userCountryCode, userPhoneNumber)
	userInn, _           = commonTypes.NewInn(userInnRaw)
	userSnils, _         = commonTypes.NewSnils(userSnilsRaw)
	userPhoto, _         = commonTypes.NewPhoto(userPhotoRaw)
	userCard, _          = commonTypes.NewCard(userCardNumber, userCardExpires)
	userRating, _        = commonTypes.NewRating(userPositiveRating, userNegativeRating)
	userState            = commonTypes.UserState(userStateRaw)

	userAggregate, _ = user.NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		userIsRemoved,
		userVersion)

	userNilAggregate, _ = user.NewUserNil(userID)

	UserAggregate    user.User
	UserNilAggregate user.Nil
)

func init() {
	userAggregate.ChangeName(userName)
	userAggregate.ChangePhone(userPhone)
	userAggregate.ChangePassport(passportEntity)
	userAggregate.ChangeDrivingLicense(drivingLicenseEntity)
	userAggregate.ChangeSnils(userSnils)
	userAggregate.ChangeInn(userInn)
	userAggregate.ChangePhoto(&userPhoto)
	userAggregate.ChangeCard(userCard)
	userAggregate.GrantFullAccess(*userOtherID)
	userAggregate.Activate()

	UserAggregate = *userAggregate
	UserNilAggregate = *userNilAggregate
}
