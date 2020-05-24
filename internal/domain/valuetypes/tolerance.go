package valuetypes

import (
	"errors"
	"fmt"
)

type Accessor int

const (
	AccessorContacts Accessor = iota + 1
	AccessorProfile
	AccessorPhone
	AccessorPassport
	AccessorDriverLicense
)

var ErrInvalidUserID = errors.New("invalid user ID")

func (a Accessor) String() string {
	switch a {
	case AccessorContacts:
		return "contacts"
	case AccessorProfile:
		return "profile"
	case AccessorPhone:
		return "phone"
	case AccessorPassport:
		return "passport"
	case AccessorDriverLicense:
		return "driver license"
	default:
		return ""
	}
}

type Tolerance struct {
	id        UserID
	accessors []Accessor
}

func NewTolerance(id *UserID, accessors []Accessor) (*Tolerance, error) {
	if id == nil {
		return nil, ErrInvalidUserID
	}

	return &Tolerance{id: *id, accessors: accessors}, nil
}

func (t Tolerance) AddFullAccess() Tolerance {
	accessors := []Accessor{
		AccessorContacts,
		AccessorProfile,
		AccessorPhone,
		AccessorPassport,
		AccessorDriverLicense}

	tolerance, _ := NewTolerance(&t.id, accessors)
	return *tolerance
}

func (t Tolerance) AddAccess(accessor Accessor) Tolerance {
	if t.HasAccess(accessor) {
		return t
	}

	accessors := append(t.accessors, accessor)
	tolerance, _ := NewTolerance(&t.id, accessors)
	return *tolerance
}

func (t Tolerance) RemoveAccess(accessor Accessor) Tolerance {
	if !t.HasAccess(accessor) {
		return t
	}

	var accessors []Accessor
	for _, val := range t.accessors {
		if val == accessor {
			continue
		}

		accessors = append(accessors, val)
	}

	tolerance, _ := NewTolerance(&t.id, accessors)
	return *tolerance
}

func (t Tolerance) HasAccess(accessor Accessor) bool {
	for _, val := range t.accessors {
		if val == accessor {
			return true
		}
	}
	return false
}

func (t *Tolerance) String() string {
	text := "accessors: "
	for idx, val := range t.accessors {
		if idx == len(t.accessors)-1 {
			text += val.String()
			break
		}

		text += fmt.Sprintf("%v, ", val.String())
	}

	return text
}

func (t *Tolerance) GetAccessors() []Accessor {
	if len(t.accessors) == 0 {
		return nil
	}

	accessors := make([]Accessor, len(t.accessors))
	copy(accessors, t.accessors)

	return accessors
}
