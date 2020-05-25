package validator

import (
	"errors"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"
)

type Validator interface {
	Fssp(*passport.PassportSnapshot) (bool, error)
	Mvd(*passport.PassportSnapshot) (bool, error)
	Ufms(*passport.PassportSnapshot) (bool, error)
	Gibdd(*driving_license.DrivingLicenseSnapshot) (bool, error)
}

type Client struct {
}

func (c *Client) Fssp(_ *passport.PassportSnapshot) (bool, error) {
	return false, errors.New("not implemented")
}

func (c *Client) Mvd(_ *passport.PassportSnapshot) (bool, error) {
	return false, errors.New("not implemented")
}

func (c *Client) Ufms(_ *passport.PassportSnapshot) (bool, error) {
	return false, errors.New("not implemented")
}

func (c *Client) Gibdd(_ *driving_license.DrivingLicenseSnapshot) (bool, error) {
	return false, errors.New("not implemented")
}
