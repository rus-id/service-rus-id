package valuetypes

import "errors"

var ErrInvalidResidence = errors.New("invalid address")

type DrivingLicenseResidence struct {
	value string
}

func NewResidence(value string) (*DrivingLicenseResidence, error) {
	if value == "" {
		return nil, ErrInvalidResidence
	}

	return &DrivingLicenseResidence{value: value}, nil
}
