package valuetypes

import "fmt"

type Accessor int

const (
	AccessorContacts Accessor = iota + 1
	AccessorProfile
	AccessorPhone
	AccessorPassport
	AccessorDriverLicense
)

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

// TODO: Может переделать на map[Accessor]struct{}
type Tolerance struct {
	id        UserID
	accessors []Accessor
}

func NewTolerance(id UserID, accessors []Accessor) *Tolerance {
	return &Tolerance{id: id, accessors: accessors}
}

func (t *Tolerance) AddFullAccess() *Tolerance {
	accessors := []Accessor{
		AccessorContacts,
		AccessorProfile,
		AccessorPhone,
		AccessorPassport,
		AccessorDriverLicense}

	return NewTolerance(t.id, accessors)
}

func (t *Tolerance) AddAccess(accessor Accessor) *Tolerance {
	if t.HasAccess(accessor) {
		return t
	}

	accessors := append(t.accessors, accessor)
	return NewTolerance(t.id, accessors)
}

func (t *Tolerance) RemoveAccess(accessor Accessor) *Tolerance {
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

	return NewTolerance(t.id, accessors)
}

func (t *Tolerance) HasAccess(accessor Accessor) bool {
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

		text += fmt.Sprintf(" %v, ", val.String())
	}

	return text
}
