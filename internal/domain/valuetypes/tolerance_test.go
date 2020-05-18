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
			t.Errorf("expected text: %s, act: %s", val.text, act)
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

	tolerance := NewTolerance(userID, nil)
	if act := len(tolerance.GetAccessors()); act != 0 {
		t.Errorf("accessories must be empty, act %v", act)
	}

	tolerance = tolerance.AddFullAccess()
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

	tolerance := NewTolerance(userID, nil)
	if act := len(tolerance.GetAccessors()); act != 0 {
		t.Errorf("accessories must be empty, act %v", act)
	}

	tolerance = tolerance.AddAccess(AccessorPhone)
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

	tolerance := NewTolerance(userID, nil)
	tolerance = tolerance.AddFullAccess()
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

	tolerance := NewTolerance(userID, nil)
	tolerance = tolerance.AddFullAccess()
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
