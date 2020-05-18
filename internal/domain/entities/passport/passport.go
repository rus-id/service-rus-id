package passport

import (
	"time"

	valuetypes2 "github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport/valuetypes"
	valuetypes3 "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

type Passport struct {
	id           valuetypes2.PassportID
	name         valuetypes3.Name
	birthday     time.Time
	issue        valuetypes2.PassportIssue
	registration valuetypes3.Address
	validation   valuetypes2.PassportValidation
}

func NewPassport() {
}
