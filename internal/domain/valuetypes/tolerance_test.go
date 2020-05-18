package valuetypes

import "testing"

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
		act := val.accessor.String()
		if act != val.text {
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

	if len(tolerance.accessors) != 0 {
		t.Errorf("accessories must me empty")
	}

	tolerance = tolerance.AddFullAccess()

	if len(tolerance.accessors) != 5 {
		t.Errorf("accessories must me 5")
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

	if len(tolerance.accessors) != 0 {
		t.Errorf("accessories must me empty")
	}

	tolerance = tolerance.AddAccess(AccessorPhone)
	tolerance = tolerance.AddAccess(AccessorPassport)

	if len(tolerance.accessors) != 2 {
		t.Errorf("accessories must me 2")
	}

	for _, val := range data {
		if act := tolerance.HasAccess(val.accessor); act == val.hasAccess {
			t.Errorf("expected %v, act: %v", val.hasAccess, act)
		}
	}
}
