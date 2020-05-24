package driving_license

import (
	"errors"
	"fmt"
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
	id           valuetypes.DrivingLicenseID
	category     valuetypes.DrivingLicenseCategory
	name         commonTypes.Name
	birthday     time.Time
	issued       time.Time
	expired      time.Time
	residence    valuetypes.DrivingLicenseResidence
	specialMarks string
	validation   valuetypes.DrivingLicenseValidation
}

func NewDrivingLicense(
	id *valuetypes.DrivingLicenseID,
	category valuetypes.DrivingLicenseCategory,
	name *commonTypes.Name,
	birthday *time.Time,
	issued *time.Time,
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

	if issued == nil {
		return nil, ErrInvalidIssue
	}

	if expired.Unix() <= issued.Unix() {
		return nil, ErrInvalidDates
	}

	if residence == nil {
		return nil, ErrInvalidResidence
	}

	if validation == nil {
		return nil, ErrInvalidValidation
	}

	return &DrivingLicense{
		id:           *id,
		category:     category,
		name:         *name,
		birthday:     *birthday,
		issued:       *issued,
		expired:      *expired,
		residence:    *residence,
		specialMarks: specialMarks,
		validation:   *validation,
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

	d.name = *name
}

func (d *DrivingLicense) ChangeBirthday(birthday *time.Time) {
	if birthday == nil {
		return
	}

	d.birthday = *birthday
}

func (d *DrivingLicense) ChangeIssued(issued *time.Time) {
	if issued == nil {
		return
	}

	d.issued = *issued
}

func (d *DrivingLicense) ChangeExpired(expired *time.Time) {
	if expired == nil {
		return
	}

	d.expired = *expired
}

func (d *DrivingLicense) ChangeResidence(residence *valuetypes.DrivingLicenseResidence) {
	if residence == nil {
		return
	}

	d.residence = *residence
}

func (d *DrivingLicense) ChangeSpecialMark(specialMark string) {
	d.specialMarks = specialMark
}

func (d *DrivingLicense) ChangeValidation(validation *valuetypes.DrivingLicenseValidation) {
	if validation == nil {
		return
	}

	d.validation = *validation
}

//Getters

func (d *DrivingLicense) GetID() valuetypes.DrivingLicenseID {
	return d.id
}

func (d *DrivingLicense) GetCategory() valuetypes.DrivingLicenseCategory {
	return d.category
}

func (d *DrivingLicense) GetName() commonTypes.Name {
	return d.name
}

func (d *DrivingLicense) GetBirthday() time.Time {
	return d.birthday
}

func (d *DrivingLicense) GetIssued() time.Time {
	return d.issued
}

func (d *DrivingLicense) GetExpired() time.Time {
	return d.expired
}

func (d *DrivingLicense) GetResidence() valuetypes.DrivingLicenseResidence {
	return d.residence
}

func (d *DrivingLicense) GetSpecialMarks() string {
	return d.specialMarks
}

func (d *DrivingLicense) GetValidation() valuetypes.DrivingLicenseValidation {
	return d.validation
}

func (d *DrivingLicense) String() string {
	return fmt.Sprintf("ID: %v. Category %v. Name: %v. Birthday: %v. Issued: %v. Exppired: %v. Residence: %v. Special Marks: %v. Validation: %v",
		d.id,
		d.category,
		d.name,
		d.birthday.Format("02.01.2006"),
		d.issued.Format("02.01.2006"),
		d.expired.Format("02.01.2006"),
		d.residence,
		d.specialMarks,
		d.validation)
}
