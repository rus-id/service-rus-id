package driving_license

import (
	"errors"
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license/valuetypes"
	commonTypes "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

var (
	ErrInvalidID         = errors.New("driving license ID not specified")
	ErrInvalidCategory   = errors.New("driving license invalid category")
	ErrInvalidName       = errors.New("driving license name not specified")
	ErrInvalidBirthday   = errors.New("driving license invalid birthday")
	ErrInvalidIssue      = errors.New("driving license invalid issued")
	ErrInvalidExpired    = errors.New("driving license invalid expired")
	ErrInvalidDates      = errors.New("driving license expired date must be greater then issued date")
	ErrInvalidResidence  = errors.New("driving license residence not specified")
	ErrInvalidValidation = errors.New("driving license validation not specified")
)

type DrivingLicense struct {
	id           *valuetypes.DrivingLicenseID
	category     valuetypes.DrivingLicenseCategory
	name         *commonTypes.Name
	birthday     *time.Time
	issued       *time.Time
	expired      *time.Time
	residence    *valuetypes.DrivingLicenseResidence
	specialMarks string
	validation   *valuetypes.DrivingLicenseValidation
}

func NewDrivingLicense(
	id *valuetypes.DrivingLicenseID,
	category valuetypes.DrivingLicenseCategory,
	name *commonTypes.Name,
	birthday *time.Time,
	issue *time.Time,
	expired *time.Time,
	residence *valuetypes.DrivingLicenseResidence,
	specialMarks string,
	validation *valuetypes.DrivingLicenseValidation) (*DrivingLicense, error) {
	if id == nil {
		return nil, ErrInvalidID
	}

	if !category.IsValid() {
		return nil, ErrInvalidCategory
	}

	if name == nil {
		return nil, ErrInvalidName
	}

	if birthday == nil {
		return nil, ErrInvalidBirthday
	}

	if expired == nil {
		return nil, ErrInvalidExpired
	}

	if issue == nil {
		return nil, ErrInvalidIssue
	}

	if expired.Unix() <= issue.Unix() {
		return nil, ErrInvalidDates
	}

	if residence == nil {
		return nil, ErrInvalidResidence
	}

	if validation == nil {
		return nil, ErrInvalidValidation
	}

	return &DrivingLicense{
		id:           id,
		category:     category,
		name:         name,
		birthday:     birthday,
		issued:       issue,
		expired:      expired,
		residence:    residence,
		specialMarks: specialMarks,
	}, nil
}

// Setters

func (d *DrivingLicense) ChangeCategory(category valuetypes.DrivingLicenseCategory) {
	if !category.IsValid() {
		return
	}

	d.category = category
}

func (d *DrivingLicense) ChangeName(name *commonTypes.Name) {
	if name == nil {
		return
	}

	d.name = name
}

func (d *DrivingLicense) ChangeBirthday(birthday *time.Time) {
	if birthday == nil {
		return
	}

	d.birthday = birthday
}

func (d *DrivingLicense) ChangeIssued(issued *time.Time) {
	if issued == nil {
		return
	}

	d.issued = issued
}

func (d *DrivingLicense) ChangeExpired(expired *time.Time) {
	if expired == nil {
		return
	}

	d.expired = expired
}

func (d *DrivingLicense) ChangeResidence(residence *valuetypes.DrivingLicenseResidence) {
	if residence == nil {
		return
	}

	d.residence = residence
}

func (d *DrivingLicense) ChangeSpecialMark(specialMark string) {
	d.specialMarks = specialMark
}

func (d *DrivingLicense) ChangeValidation(validation *valuetypes.DrivingLicenseValidation) {
	if validation == nil {
		return
	}

	d.validation = validation
}

//Getters

func (d *DrivingLicense) GetID() *valuetypes.DrivingLicenseID {
	return d.id
}

func (d *DrivingLicense) GetCategory() valuetypes.DrivingLicenseCategory {
	return d.category
}

func (d *DrivingLicense) GetName() *commonTypes.Name {
	return d.name
}

func (d *DrivingLicense) GetBirthday() *time.Time {
	return d.birthday
}

func (d *DrivingLicense) GetIssued() *time.Time {
	return d.issued
}

func (d *DrivingLicense) GetExpired() *time.Time {
	return d.expired
}

func (d *DrivingLicense) GetResidence() *valuetypes.DrivingLicenseResidence {
	return d.residence
}

func (d *DrivingLicense) GetSpecialMarks() string {
	return d.specialMarks
}

func (d *DrivingLicense) GetValidation() *valuetypes.DrivingLicenseValidation {
	return d.validation
}
