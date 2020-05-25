package driving_license_test

import (
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license/valuetypes"
	commonTypes "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

const (
	serial             = "7777"
	number             = "777777"
	firstName          = "Boris"
	lastName           = "Goldovsky"
	country            = "Russia"
	gibddValidation    = true
	documentValidation = false
	specialMarks       = "empty mark"
)

var (
	birthday      = time.Date(1986, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	issued        = time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	expires       = time.Date(2025, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	birthdayStamp = birthday.Unix()
	issuedStamp   = issued.Unix()
	expiresStamp  = expires.Unix()

	middleName   *string = nil
	id, _                = valuetypes.NewDrivingLicenseID(serial, number)
	category             = valuetypes.DrivingLicenseA
	name, _              = commonTypes.NewName(firstName, middleName, lastName)
	residence, _         = valuetypes.NewResidence(country)
	validation           = valuetypes.NewDrivingLicenseValidation(gibddValidation, documentValidation)
)
