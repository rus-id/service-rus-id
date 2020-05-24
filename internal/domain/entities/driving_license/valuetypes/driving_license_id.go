package valuetypes

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrInvalidDrivingLicenseSerial = errors.New("invalid driving license serial identity")
	ErrInvalidDrivingLicenseNumber = errors.New("invalid driving license number number")

	dlSerialRegexp = regexp.MustCompile("^[0-9]{4}$")
	dlNumberRegexp = regexp.MustCompile("^[0-9]{6}$")
)

type DrivingLicenseID struct {
	serial string
	number string
}

func NewDrivingLicenseID(serial, number string) (*DrivingLicenseID, error) {
	if !dlSerialRegexp.MatchString(serial) {
		return nil, ErrInvalidDrivingLicenseSerial
	}

	if !dlNumberRegexp.MatchString(number) {
		return nil, ErrInvalidDrivingLicenseNumber
	}

	return &DrivingLicenseID{serial: serial, number: number}, nil
}

func (id DrivingLicenseID) GetSerial() string {
	return id.serial
}

func (id DrivingLicenseID) GetNumber() string {
	return id.number
}

func (id DrivingLicenseID) String() string {
	serialRunes := []rune(id.serial)

	return fmt.Sprintf("%v %v\t%v",
		string(serialRunes[0:2]),
		string(serialRunes[2:4]),
		id.number)
}
