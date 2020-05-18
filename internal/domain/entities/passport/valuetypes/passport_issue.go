package valuetypes

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrInvalidOrganisation = errors.New("invalid organisation")
	ErrInvalidCode         = errors.New("invalid organisation")
)

type PassportIssue struct {
	organisation string
	date         time.Time
	code         string
}

func NewPassportIssue(organisation string, date time.Time, code string) (*PassportIssue, error) {
	if organisation == "" {
		return nil, ErrInvalidOrganisation
	}

	if code == "" {
		return nil, ErrInvalidCode
	}

	return &PassportIssue{organisation: organisation, date: date, code: code}, nil
}

func (pi *PassportIssue) GetOrganisation() string {
	return pi.organisation
}

func (pi *PassportIssue) GetDate() time.Time {
	return pi.date
}

func (pi *PassportIssue) GetCode() string {
	return pi.code
}

func (pi *PassportIssue) String() string {
	return fmt.Sprintf("Issued by %s, %v, code %v",
		pi.organisation,
		pi.date.Format("02.01.2006"),
		pi.code)
}
