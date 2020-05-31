package user_test

import (
	"testing"

	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/user"
)

func TestNewUserNil(t *testing.T) {
	user, err := NewUserNil(userID)

	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	if id := user.GetID(); id != *userID {
		t.Errorf("expected: %v, act: %v", *userID, id)
	}

	if act := user.IsRemoved(); !act {
		t.Errorf("expected: %v, act: %v", true, act)
	}
}

func TestNewUserNil_Error(t *testing.T) {
	user, err := NewUserNil(nil)

	if err != ErrInvalidID {
		t.Errorf("expected: %v, act: %v", ErrInvalidID, err)
	}

	if user != nil {
		t.Errorf("expected: %v, act: %v", nil, user)
	}
}

func TestUserNil_Methods(t *testing.T) {
	user, err := NewUserNil(userID)

	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	if id := user.GetID(); id != *userID {
		t.Errorf("expected: %v, act: %v", *userID, id)
	}

	user.ChangeName(userName)
	user.ChangePhone(userPhone)
	user.ChangePassport(passportEntity)
	user.ChangeDrivingLicense(drivingLicenseEntity)
	user.ChangeSnils(userSnils)
	user.ChangeInn(userInn)
	user.ChangePhoto(&userPhoto)
	user.ChangeCard(userCard)
	user.IncreaseRating()
	user.DecreaseRating()
	user.GrantAccess(*userOtherID, valuetypes.AccessorContacts)
	user.RevokeAccess(*userOtherID, valuetypes.AccessorContacts)
	user.GrantFullAccess(*userOtherID)
	user.RevokeFullAccess(*userOtherID)
	user.Block()
	user.Activate()
	user.Remove()

	if id := user.GetID(); id != *userID {
		t.Errorf("expected: %v, act: %v", *userID, id)
	}
}

func TestUserNil_String(t *testing.T) {
	expected := "User Aggregate ID 059b4e12-6983-4806-bd5a-cc3433e78f66\nRemoved: true\n"

	user, _ := NewUserNil(userID)

	if act := user.String(); act != expected {
		t.Errorf("expected: %v, act: %v", expected, act)
	}
}
