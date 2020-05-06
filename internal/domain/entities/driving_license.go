package entities

import (
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/valuetypes"
)

type DrivingLicense struct {
	id           valuetypes.DrivingLicenseID
	category     valuetypes.DrivingLicenseCategory
	name         valuetypes.Name
	birthday     time.Time
	issue        time.Time
	expired      time.Time
	residence    valuetypes.DrivingLicenseResidence
	specialMarks string
}
