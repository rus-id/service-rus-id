package valuetypes

import (
	"errors"
	"fmt"
	"regexp"
)

const RusCode = 7

var (
	ErrInvalidCountryCode = errors.New("invalid country code")
	ErrInvalidPhoneNumber = errors.New("invalid phone number code")
)

var phoneRegexp = regexp.MustCompile("^[0-9]{10}$")

type Phone struct {
	Code   int
	Number string
}

func NewPhone(code int, number string) (*Phone, error) {
	if ok, err := validateNumber(code, number); !ok {
		return nil, err
	}
	return &Phone{
		Code:   code,
		Number: number,
	}, nil
}

func (p *Phone) String() string {
	return fmt.Sprintf("+%v%v", p.Code, p.Number)
}

func validateNumber(code int, number string) (bool, error) {
	if code != RusCode {
		return false, ErrInvalidCountryCode
	}

	if !phoneRegexp.MatchString(number) {
		return false, ErrInvalidPhoneNumber
	}

	return true, nil
}
