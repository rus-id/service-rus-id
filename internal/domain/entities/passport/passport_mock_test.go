package passport_test

import (
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport/valuetypes"
	commonTypes "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

const (
	serial             = "7777"
	number             = "777777"
	firstName          = "Boris"
	lastName           = "Goldovsky"
	ufmsValidation     = true
	mvdValidation      = false
	fsspValidation     = true
	address            = "Russia, Moscow"
	issuedOrganisation = "MVD"
	issuedCode         = "770-77"
	documentValidation = false
)

var (
	birthday      = time.Date(1986, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	issuedDate    = time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	birthdayStamp = birthday.Unix()
	issuedStamp   = issuedDate.Unix()

	middleName      *string = nil
	id, _                   = valuetypes.NewPassportID(serial, number)
	name, _                 = commonTypes.NewName(firstName, nil, lastName)
	issued, _               = valuetypes.NewPassportIssue(issuedOrganisation, issuedDate, issuedCode)
	registration, _         = commonTypes.NewAddress(address)
	validation              = valuetypes.NewPassportValidation(ufmsValidation, mvdValidation, fsspValidation, documentValidation)
)
