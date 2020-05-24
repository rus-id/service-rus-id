package valuetypes

import (
	"errors"
	"fmt"
	"regexp"
)

const RusCountryCode CountryCode = 7

type CountryCode int

var (
	ErrInvalidCountryCode = errors.New("invalid country code")
	ErrInvalidPhoneNumber = errors.New("invalid phone number code")
)

var phoneRegexp = regexp.MustCompile("^[0-9]{10}$")

type Phone struct {
	code   CountryCode
	number string
}

func NewPhone(code CountryCode, number string) (*Phone, error) {
	if ok, err := validateNumber(code, number); !ok {
		return nil, err
	}

	return &Phone{
		code:   code,
		number: number,
	}, nil
}

func (p Phone) GetCode() CountryCode {
	return p.code
}

func (p Phone) GetNumber() string {
	return p.number
}

func (p Phone) String() string {
	runes := []rune(p.number)

	return fmt.Sprintf("+%v(%v)%v-%v-%v",
		p.code,
		string(runes[:3]),
		string(runes[3:6]),
		string(runes[6:8]),
		string(runes[8:10]))
}

func validateNumber(code CountryCode, number string) (bool, error) {
	if code != RusCountryCode {
		return false, ErrInvalidCountryCode
	}

	if !phoneRegexp.MatchString(number) {
		return false, ErrInvalidPhoneNumber
	}

	return true, nil
}
