package entities

import "github.com/bgoldovsky/service-rus-id/internal/domain/entities/valuetypes"

type DrivingLicense struct {
	id   valuetypes.LicenseID
	name valuetypes.Name
}
