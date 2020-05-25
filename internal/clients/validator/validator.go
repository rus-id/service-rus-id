package validator

import (
	"errors"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"
)

type Validator interface {
	Fssp(*passport.Snapshot) (bool, error)
	Mvd(*passport.Snapshot) (bool, error)
	Ufms(*passport.Snapshot) (bool, error)
	Gibdd(*driving_license.Snapshot) (bool, error)
}

type Client struct {
}

func (c *Client) Fssp(_ *passport.Snapshot) (bool, error) {
	return false, errors.New("not implemented")
}

func (c *Client) Mvd(_ *passport.Snapshot) (bool, error) {
	return false, errors.New("not implemented")
}

func (c *Client) Ufms(_ *passport.Snapshot) (bool, error) {
	return false, errors.New("not implemented")
}

func (c *Client) Gibdd(_ *driving_license.Snapshot) (bool, error) {
	return false, errors.New("not implemented")
}
