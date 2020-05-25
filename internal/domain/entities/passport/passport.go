package passport

import (
	"errors"
	"fmt"
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport/valuetypes"
	commonTypes "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

var (
	ErrInvalidID           = errors.New("passport ID not specified")
	ErrInvalidName         = errors.New("passport name not specified")
	ErrInvalidBirthday     = errors.New("passport invalid birthday")
	ErrInvalidIssue        = errors.New("passport invalid issued")
	ErrInvalidRegistration = errors.New("passport registration not specified")
	ErrInvalidValidation   = errors.New("passport validation not specified")
)

type Passport struct {
	id           valuetypes.PassportID
	name         commonTypes.Name
	birthday     time.Time
	issued       valuetypes.PassportIssue
	registration commonTypes.Address
	validation   valuetypes.PassportValidation
}

func NewPassport(
	id *valuetypes.PassportID,
	name *commonTypes.Name,
	birthday *time.Time,
	issued *valuetypes.PassportIssue,
	registration *commonTypes.Address,
	validation *valuetypes.PassportValidation) (*Passport, error) {
	if id == nil {
		return nil, ErrInvalidID
	}

	if name == nil {
		return nil, ErrInvalidName
	}

	if birthday == nil {
		return nil, ErrInvalidBirthday
	}

	if issued == nil {
		return nil, ErrInvalidIssue
	}

	if registration == nil {
		return nil, ErrInvalidRegistration
	}

	if validation == nil {
		return nil, ErrInvalidValidation
	}

	return &Passport{
		id:           *id,
		name:         *name,
		birthday:     *birthday,
		issued:       *issued,
		registration: *registration,
		validation:   *validation,
	}, nil
}

// Setters

func (p *Passport) ChangeName(name *commonTypes.Name) {
	if name == nil {
		return
	}

	p.name = *name
}

func (p *Passport) ChangeBirthday(birthday *time.Time) {
	if birthday == nil {
		return
	}

	p.birthday = *birthday
}

func (p *Passport) ChangeIssued(issued *valuetypes.PassportIssue) {
	if issued == nil {
		return
	}

	p.issued = *issued
}

func (p *Passport) ChangeRegistration(registration *commonTypes.Address) {
	if registration == nil {
		return
	}

	p.registration = *registration
}

func (p *Passport) ChangeValidation(validation *valuetypes.PassportValidation) {
	if validation == nil {
		return
	}

	p.validation = *validation
}

// Getters

func (p *Passport) GetID() valuetypes.PassportID {
	return p.id
}

func (p *Passport) GetName() commonTypes.Name {
	return p.name
}

func (p *Passport) GetBirthday() time.Time {
	return p.birthday
}

func (p *Passport) GetIssued() valuetypes.PassportIssue {
	return p.issued
}

func (p *Passport) GetRegistration() commonTypes.Address {
	return p.registration
}

func (p *Passport) GetValidation() valuetypes.PassportValidation {
	return p.validation
}

func (p *Passport) String() string {
	return fmt.Sprintf("ID %v; name %v; birthday %v; issued %v; registration %v; %v",
		p.id,
		p.name,
		p.birthday.Format("02.01.2006"),
		p.issued,
		p.registration,
		p.validation)
}
