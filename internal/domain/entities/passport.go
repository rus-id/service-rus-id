package entities

import (
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/valuetypes"
)

type Passport struct {
	id           valuetypes.PassportID
	name         valuetypes.Name
	birthday     time.Time
	issue        valuetypes.PassportIssue
	registration valuetypes.Address
}

func NewPassport() {
}
