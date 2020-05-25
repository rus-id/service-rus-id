package valuetypes

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrInvalidPassportSerial = errors.New("invalid passport serial identity")
	ErrInvalidPassportNumber = errors.New("invalid passport number number")

	passportSerialRegexp = regexp.MustCompile("^[0-9]{4}$")
	passportNumberRegexp = regexp.MustCompile("^[0-9]{6}$")
)

type PassportID struct {
	serial string
	number string
}

func NewPassportID(serial, number string) (*PassportID, error) {
	if !passportSerialRegexp.MatchString(serial) {
		return nil, ErrInvalidPassportSerial
	}

	if !passportNumberRegexp.MatchString(number) {
		return nil, ErrInvalidPassportNumber
	}

	return &PassportID{serial: serial, number: number}, nil
}

func (id PassportID) GetSerial() string {
	return id.serial
}

func (id PassportID) GetNumber() string {
	return id.number
}

func (id PassportID) String() string {
	serialRunes := []rune(id.serial)

	return fmt.Sprintf("%v %v %v",
		string(serialRunes[0:2]),
		string(serialRunes[2:4]),
		id.number)
}
