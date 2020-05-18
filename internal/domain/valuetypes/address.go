package valuetypes

import "errors"

var ErrInvalidAddress = errors.New("invalid address")

type Address string

func NewAddress(address string) (Address, error) {
	if address == "" {
		return "", ErrInvalidAddress
	}

	return Address(address), nil
}
