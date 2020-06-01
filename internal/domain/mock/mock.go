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
	PassportSerial             = "7777"
	PassportNumber             = "777777"
	PassportFirstName          = "Boris"
	PassportLastName           = "Goldovsky"
	PassportUfmsValidation     = true
	PassportMvdValidation      = false
	PassportFsspValidation     = true
	PassportAddress            = "Russia, Moscow"
	PassportIssuedOrganisation = "MVD"
	PassportIssuedCode         = "770-77"
	PassportDocumentValidation = false
)

var (
	PassportBirthday                = time.Date(1986, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	PassportIssuedDate              = time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	PassportBirthdayStamp           = PassportBirthday.Unix()
	PassportIssuedStamp             = PassportIssuedDate.Unix()
	PassportMiddleName      *string = nil
	PassportID, _                   = passportTypes.NewPassportID(PassportSerial, PassportNumber)
	PassportName, _                 = commonTypes.NewName(PassportFirstName, nil, PassportLastName)
	PassportIssued, _               = passportTypes.NewPassportIssue(PassportIssuedOrganisation, PassportIssuedDate, PassportIssuedCode)
	PassportRegistration, _         = commonTypes.NewAddress(PassportAddress)
	PassportValidation              = passportTypes.NewPassportValidation(PassportUfmsValidation, PassportMvdValidation, PassportFsspValidation, PassportDocumentValidation)

	PassportEntity, _ = passport.NewPassport(
		PassportID,
		PassportName,
		&PassportBirthday,
		PassportIssued,
		PassportRegistration,
		PassportValidation)

	PassportSnapshot = passport.NewSnapshot(
		PassportSerial,
		PassportNumber,
		PassportFirstName,
		PassportMiddleName,
		PassportLastName,
		PassportBirthdayStamp,
		PassportIssuedOrganisation,
		PassportIssuedStamp,
		PassportIssuedCode,
		PassportAddress,
		PassportUfmsValidation,
		PassportMvdValidation,
		PassportFsspValidation,
		PassportDocumentValidation)
)

// Driving License

const (
	DLSerial             = "7777"
	DLNumber             = "777777"
	DLFirstName          = "Boris"
	DLLastName           = "Goldovsky"
	DLCountry            = "Russia"
	DLGibddValidation    = true
	DLDocumentValidation = false
	DLSpecialMarks       = "empty mark"
)

var (
	DLBirthday              = time.Date(1986, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	DLIssued                = time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	DLExpires               = time.Date(2025, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	DLBirthdayStamp         = DLBirthday.Unix()
	DLIssuedStamp           = DLIssued.Unix()
	DLExpiresStamp          = DLExpires.Unix()
	DLMiddleName    *string = nil
	DLId, _                 = dlValues.NewDrivingLicenseID(DLSerial, DLNumber)
	DLCategory              = dlValues.DrivingLicenseA
	DLName, _               = commonTypes.NewName(DLFirstName, DLMiddleName, DLLastName)
	DLResidence, _          = dlValues.NewResidence(DLCountry)
	DLValidation            = dlValues.NewDrivingLicenseValidation(DLGibddValidation, DLDocumentValidation)

	DrivingLicenseEntity, _ = driving_license.NewDrivingLicense(
		DLId,
		DLCategory,
		DLName,
		&DLBirthday,
		&DLIssued,
		&DLExpires,
		DLResidence,
		DLSpecialMarks,
		DLValidation)

	DrivingLicenseSnapshot = driving_license.NewSnapshot(
		DLSerial,
		DLNumber,
		int64(DLCategory),
		DLFirstName,
		DLMiddleName,
		DLLastName,
		DLBirthdayStamp,
		DLIssuedStamp,
		DLExpiresStamp,
		DLCountry,
		DLSpecialMarks,
		DLGibddValidation,
		DLDocumentValidation)
)

// User

const (
	UserIDRaw          = "059b4e12-6983-4806-bd5a-cc3433e78f66"
	UserOtherIDRaw     = "060b4e12-6983-4806-bd5a-cc3433e78f66"
	UserFirstName      = "Boris"
	UserLastName       = "Goldovsky"
	UserCountryCode    = 7
	UserPhoneNumber    = "9039615322"
	UserNegativeRating = 40
	UserPositiveRating = 50
	UserStateRaw       = 1
	UserIsRemoved      = false
	UserVersion        = 12345
)

var (
	UserCardExpires                   = time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	UserRegistrationDate              = time.Date(2019, time.Month(3), 9, 1, 10, 30, 0, time.UTC)
	UserSnapshotDate                  = time.Date(2020, time.Month(2), 4, 1, 10, 30, 0, time.UTC)
	UserCardExpiresStamp              = UserCardExpires.Unix()
	UserRegistrationDateStamp         = UserRegistrationDate.Unix()
	UserSnapshotDateStamp             = UserSnapshotDate.Unix()
	UserMiddleName            *string = nil
	UserInnRaw                        = "926902267890"
	UserSnilsRaw                      = "59650418527"
	UserCardNumber                    = "4444333322221111"
	UserPhotoRaw                      = []byte{10, 20, 30, 40, 50}
	UserID, _                         = commonTypes.NewUserID(UserIDRaw)
	UserOtherID, _                    = commonTypes.NewUserID(UserOtherIDRaw)
	UserName, _                       = commonTypes.NewName(DLFirstName, DLMiddleName, DLLastName)
	UserPhone, _                      = commonTypes.NewPhone(UserCountryCode, UserPhoneNumber)
	UserInn, _                        = commonTypes.NewInn(UserInnRaw)
	UserSnils, _                      = commonTypes.NewSnils(UserSnilsRaw)
	UserPhoto, _                      = commonTypes.NewPhoto(UserPhotoRaw)
	UserCard, _                       = commonTypes.NewCard(UserCardNumber, UserCardExpires)
	UserRating, _                     = commonTypes.NewRating(UserPositiveRating, UserNegativeRating)
	UserState                         = commonTypes.UserState(UserStateRaw)
	UserAccessors                     = []commonTypes.Accessor{
		commonTypes.AccessorContacts,
		commonTypes.AccessorProfile,
		commonTypes.AccessorPhone,
		commonTypes.AccessorPassport,
		commonTypes.AccessorDriverLicense}
	UserTolerance, _            = commonTypes.NewTolerance(UserOtherID, UserAccessors)
	UserTolerances              = map[commonTypes.UserID]commonTypes.Tolerance{*UserOtherID: *UserTolerance}
	UserToleranceSnapshots      = user.GetToleranceSnapshot(*UserID, UserTolerances)
	UserToleranceEmptySnapshots = user.GetToleranceSnapshot(*UserID, make(map[commonTypes.UserID]commonTypes.Tolerance))

	userAggregate, _ = user.NewUser(
		UserID,
		UserName,
		UserPhone,
		&UserRegistrationDate,
		UserRating,
		UserState,
		UserIsRemoved,
		UserVersion)

	userNilAggregate, _ = user.NewUserNil(UserID)

	UserAggregate     user.User
	UserNilAggregate  user.Nil
	UserSnapshot      user.Snapshot
	UserOtherSnapshot user.Snapshot
)

func init() {
	userAggregate.ChangeName(UserName)
	userAggregate.ChangePhone(UserPhone)
	userAggregate.ChangePassport(PassportEntity)
	userAggregate.ChangeDrivingLicense(DrivingLicenseEntity)
	userAggregate.ChangeSnils(UserSnils)
	userAggregate.ChangeInn(UserInn)
	userAggregate.ChangePhoto(&UserPhoto)
	userAggregate.ChangeCard(UserCard)
	userAggregate.GrantFullAccess(*UserOtherID)
	userAggregate.Activate()

	UserAggregate = *userAggregate
	UserNilAggregate = *userNilAggregate

	userSnapshot, _ := user.GetSnapshot(userAggregate, UserSnapshotDate)
	UserSnapshot = *userSnapshot
	UserOtherSnapshot = *userSnapshot
	UserOtherSnapshot.UserID = UserOtherIDRaw

}
