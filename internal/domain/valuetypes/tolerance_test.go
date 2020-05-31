package valuetypes_test

import (
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestAccessor_String(t *testing.T) {
	data := []struct {
		accessor Accessor
		text     string
	}{
		{AccessorContacts, "contacts"},
		{AccessorProfile, "profile"},
		{AccessorPhone, "phone"},
		{AccessorPassport, "passport"},
		{AccessorDriverLicense, "driver license"},
	}

	for _, val := range data {
		if act := val.accessor.String(); act != val.text {
			t.Errorf("expected: %q, act: %q", val.text, act)
		}
	}
}

func TestNewTolerance(t *testing.T) {
	data := []struct {
		userID    *UserID
		accessors []Accessor
		err       error
	}{
		{CreateUserID(), []Accessor{AccessorContacts}, nil},
		{CreateUserID(), nil, nil},
		{nil, []Accessor{AccessorContacts}, ErrInvalidUserID},
		{nil, nil, ErrInvalidUserID},
	}

	for _, val := range data {
		act, err := NewTolerance(val.userID, val.accessors)

		if err != val.err {
			t.Errorf("expected err: %v, act: %v", val.err, err)
		}

		if err != nil {
			continue
		}

		if len(act.GetAccessors()) != len(val.accessors) {
			t.Errorf("invalid accessors")
		}
	}
}

func TestTolerance_AddFullAccess(t *testing.T) {
	userID := CreateUserID()
	accessors := []Accessor{
		AccessorContacts,
		AccessorProfile,
		AccessorPhone,
		AccessorPassport,
		AccessorDriverLicense}

	newTolerance, _ := NewTolerance(userID, nil)
	if act := len(newTolerance.GetAccessors()); act != 0 {
		t.Errorf("accessories must be empty, act %v", act)
	}

	tolerance := newTolerance.AddFullAccess()
	if act := len(tolerance.GetAccessors()); act != 5 {
		t.Errorf("accessories must be 5, act %v", act)
	}

	for _, val := range accessors {
		if act := tolerance.HasAccess(val); !act {
			t.Errorf("tolerance must be full access")
		}
	}
}

func TestTolerance_AddAccess(t *testing.T) {
	userID := CreateUserID()
	data := []struct {
		accessor  Accessor
		hasAccess bool
	}{
		{AccessorContacts, false},
		{AccessorProfile, false},
		{AccessorPhone, true},
		{AccessorPassport, true},
		{AccessorDriverLicense, false},
	}

	newTolerance, _ := NewTolerance(userID, nil)
	if act := len(newTolerance.GetAccessors()); act != 0 {
		t.Errorf("accessories must be empty, act %v", act)
	}

	tolerance := newTolerance.AddAccess(AccessorPhone)
	tolerance = tolerance.AddAccess(AccessorPassport)

	if act := len(tolerance.GetAccessors()); act != 2 {
		t.Errorf("accessories must be 2, act %v", act)
	}

	for _, val := range data {
		if act := tolerance.HasAccess(val.accessor); act != val.hasAccess {
			t.Errorf("expected %v, act: %v", val.hasAccess, act)
		}
	}
}

func TestTolerance_RemoveAccess(t *testing.T) {
	userID := CreateUserID()
	data := []struct {
		accessor  Accessor
		hasAccess bool
	}{
		{AccessorContacts, false},
		{AccessorProfile, false},
		{AccessorPhone, true},
		{AccessorPassport, true},
		{AccessorDriverLicense, true},
	}

	newTolerance, _ := NewTolerance(userID, nil)
	tolerance := newTolerance.AddFullAccess()
	if act := len(tolerance.GetAccessors()); act != 5 {
		t.Errorf("accessories must me 5, act %v", act)
	}

	tolerance = tolerance.RemoveAccess(AccessorContacts)
	tolerance = tolerance.RemoveAccess(AccessorProfile)

	if act := len(tolerance.GetAccessors()); act != 3 {
		t.Errorf("accessories must be 3, act %v", act)
	}

	for _, val := range data {
		if act := tolerance.HasAccess(val.accessor); act != val.hasAccess {
			t.Errorf("expected %v, act: %v", val.hasAccess, act)
		}
	}
}

func TestTolerance_Idempotent(t *testing.T) {
	userID := CreateUserID()
	data := []struct {
		accessor  Accessor
		hasAccess bool
	}{
		{AccessorContacts, false},
		{AccessorProfile, true},
		{AccessorPhone, true},
		{AccessorPassport, true},
		{AccessorDriverLicense, true},
	}

	newTolerance, _ := NewTolerance(userID, nil)
	tolerance := newTolerance.AddFullAccess()
	if act := len(tolerance.GetAccessors()); act != 5 {
		t.Errorf("accessories must me 5, act %v", act)
	}

	tolerance = tolerance.AddAccess(AccessorContacts)
	tolerance = tolerance.AddAccess(AccessorContacts)
	if act := len(tolerance.GetAccessors()); act != 5 {
		t.Errorf("accessories must me 5, act %v", act)
	}

	tolerance = tolerance.RemoveAccess(AccessorContacts)
	tolerance = tolerance.RemoveAccess(AccessorContacts)

	if act := len(tolerance.GetAccessors()); act != 4 {
		t.Errorf("accessories must be 4, act %v", act)
	}

	for _, val := range data {
		if act := tolerance.HasAccess(val.accessor); act != val.hasAccess {
			t.Errorf("expected %v, act: %v", val.hasAccess, act)
		}
	}
}

func TestTolerance_AnyAccessors(t *testing.T) {
	userID, _ := NewUserID("059b4e12-6983-4806-bd5a-cc3433e78f66")
	accessors := []Accessor{AccessorContacts}

	data := []struct {
		accessors []Accessor
		exp       bool
	}{
		{accessors, true},
		{[]Accessor{}, false},
	}

	for _, val := range data {
		tolerance, _ := NewTolerance(userID, val.accessors)
		if act := tolerance.AnyAccessors(); act != val.exp {
			t.Errorf("expexted %v, act %v", val.exp, act)
		}
	}
}

func TestTolerance_String(t *testing.T) {
	userID, _ := NewUserID("059b4e12-6983-4806-bd5a-cc3433e78f66")
	accessors := []Accessor{
		AccessorContacts,
		AccessorProfile,
		AccessorPhone,
		AccessorPassport,
		AccessorDriverLicense}

	data := []struct {
		accessors []Accessor
		exp       string
	}{
		{accessors, "contacts, profile, phone, passport, driver license"},
		{[]Accessor{}, ""},
	}

	for _, val := range data {
		tolerance, _ := NewTolerance(userID, val.accessors)
		if act := tolerance.String(); act != val.exp {
			t.Errorf("expexted %q, act %q", val.exp, act)
		}
	}
}
