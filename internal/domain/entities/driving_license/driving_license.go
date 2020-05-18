package driving_license

import (
	"time"

	valuetypes2 "github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license/valuetypes"
	valuetypes3 "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

type DrivingLicense struct {
	id           valuetypes2.DrivingLicenseID
	category     valuetypes2.DrivingLicenseCategory
	name         valuetypes3.Name
	birthday     time.Time
	issue        time.Time
	expired      time.Time
	residence    valuetypes2.DrivingLicenseResidence
	specialMarks string
}
