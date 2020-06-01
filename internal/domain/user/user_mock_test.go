package user_test

/*
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
	passportBirthday                = time.Date(1986, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	passportIssuedDate              = time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	passportBirthdayStamp           = passportBirthday.Unix()
	passportIssuedStamp             = passportIssuedDate.Unix()
	passportMiddleName      *string = nil
	passportID, _                   = passportTypes.NewPassportID(passportSerial, passportNumber)
	passportName, _                 = commonTypes.NewName(passportFirstName, nil, passportLastName)
	passportIssued, _               = passportTypes.NewPassportIssue(passportIssuedOrganisation, passportIssuedDate, passportIssuedCode)
	passportRegistration, _         = commonTypes.NewAddress(passportAddress)
	passportValidation              = passportTypes.NewPassportValidation(passportUfmsValidation, passportMvdValidation, passportFsspValidation, passportDocumentValidation)

	passportEntity, _ = passport.NewPassport(
		passportID,
		passportName,
		&passportBirthday,
		passportIssued,
		passportRegistration,
		passportValidation)

	passportSnapshot = passport.NewSnapshot(
		passportSerial,
		passportNumber,
		passportFirstName,
		passportMiddleName,
		passportLastName,
		passportBirthdayStamp,
		passportIssuedOrganisation,
		passportIssuedStamp,
		passportIssuedCode,
		passportAddress,
		passportUfmsValidation,
		passportMvdValidation,
		passportFsspValidation,
		passportDocumentValidation)
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
	dlBirthday              = time.Date(1986, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	dlIssued                = time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	dlExpires               = time.Date(2025, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	dlBirthdayStamp         = dlBirthday.Unix()
	dlIssuedStamp           = dlIssued.Unix()
	dlExpiresStamp          = dlExpires.Unix()
	dlMiddleName    *string = nil
	dlID, _                 = dlValues.NewDrivingLicenseID(dlSerial, dlNumber)
	dlCategory              = dlValues.DrivingLicenseA
	dlName, _               = commonTypes.NewName(dlFirstName, dlMiddleName, dlLastName)
	dlResidence, _          = dlValues.NewResidence(dlCountry)
	dlValidation            = dlValues.NewDrivingLicenseValidation(dlGibddValidation, dlDocumentValidation)

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

	drivingLicenseSnapshot = driving_license.NewSnapshot(
		dlSerial,
		dlNumber,
		int64(dlCategory),
		dlFirstName,
		dlMiddleName,
		dlLastName,
		dlBirthdayStamp,
		dlIssuedStamp,
		dlExpiresStamp,
		dlCountry,
		dlSpecialMarks,
		dlGibddValidation,
		dlDocumentValidation)
)

// User

const (
	userIDRaw          = "059b4e12-6983-4806-bd5a-cc3433e78f66"
	userOtherIDRaw     = "059b4e12-6983-4806-bd5a-cc3433e78f66"
	userFirstName      = "Boris"
	userLastName       = "Goldovsky"
	userCountryCode    = 7
	userPhoneNumber    = "9039615322"
	userNegativeRating = 40
	userPositiveRating = 50
	userStateRaw       = 1
	userIsRemoved      = true
	userVersion        = 12345
)

var (
	userCardExpires                   = time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	userRegistrationDate              = time.Date(2019, time.Month(3), 9, 1, 10, 30, 0, time.UTC)
	userSnapshotDate                  = time.Date(2020, time.Month(2), 4, 1, 10, 30, 0, time.UTC)
	userCardExpiresStamp              = userCardExpires.Unix()
	userRegistrationDateStamp         = userRegistrationDate.Unix()
	userSnapshotDateStamp             = userSnapshotDate.Unix()
	userMiddleName            *string = nil
	userInnRaw                        = "926902267890"
	userSnilsRaw                      = "59650418527"
	userCardNumber                    = "4444333322221111"
	userPhotoRaw                      = []byte{10, 20, 30, 40, 50}
	userID, _                         = commonTypes.NewUserID(userIDRaw)
	userOtherID, _                    = commonTypes.NewUserID(userOtherIDRaw)
	userName, _                       = commonTypes.NewName(dlFirstName, dlMiddleName, dlLastName)
	userPhone, _                      = commonTypes.NewPhone(userCountryCode, userPhoneNumber)
	userInn, _                        = commonTypes.NewInn(userInnRaw)
	userSnils, _                      = commonTypes.NewSnils(userSnilsRaw)
	userPhoto, _                      = commonTypes.NewPhoto(userPhotoRaw)
	userCard, _                       = commonTypes.NewCard(userCardNumber, userCardExpires)
	userRating, _                     = commonTypes.NewRating(userPositiveRating, userNegativeRating)
	userState                         = commonTypes.UserState(userStateRaw)
	userAccessors                     = []commonTypes.Accessor{
		commonTypes.AccessorContacts,
		commonTypes.AccessorProfile,
		commonTypes.AccessorPhone,
		commonTypes.AccessorPassport,
		commonTypes.AccessorDriverLicense}
	userTolerance, _            = commonTypes.NewTolerance(userOtherID, userAccessors)
	userTolerances              = map[commonTypes.UserID]commonTypes.Tolerance{*userOtherID: *userTolerance}
	userToleranceSnapshots      = user.GetToleranceSnapshot(*userID, userTolerances)
	userToleranceEmptySnapshots = user.GetToleranceSnapshot(*userID, make(map[commonTypes.UserID]commonTypes.Tolerance))
)

*/
