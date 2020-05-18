package driving_license

import (
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license/valuetypes"
	commonTypes "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

type DrivingLicense struct {
	id           valuetypes.DrivingLicenseID
	category     valuetypes.DrivingLicenseCategory
	name         commonTypes.Name
	birthday     time.Time
	issue        time.Time
	expired      time.Time
	residence    valuetypes.DrivingLicenseResidence
	specialMarks string
}

func NewDrivingLicense(
	id valuetypes.DrivingLicenseID,
	category valuetypes.DrivingLicenseCategory,
	name commonTypes.Name,
	birthday time.Time,
	issue time.Time,
	expired time.Time,
	residence valuetypes.DrivingLicenseResidence,
	specialMarks string) *DrivingLicense {
	return &DrivingLicense{
		id:           id,
		category:     category,
		name:         name,
		birthday:     birthday,
		issue:        issue,
		expired:      expired,
		residence:    residence,
		specialMarks: specialMarks,
	}
}
