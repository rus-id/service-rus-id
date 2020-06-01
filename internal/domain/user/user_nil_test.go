package user_test

import (
	"testing"

	"github.com/bgoldovsky/service-rus-id/internal/domain/mock"

	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/user"
)

func TestNewUserNil(t *testing.T) {
	user, err := NewUserNil(mock.UserID)

	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	if id := user.GetID(); id != *mock.UserID {
		t.Errorf("expected: %v, act: %v", *mock.UserID, id)
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
	user, err := NewUserNil(mock.UserID)

	if err != nil {
		t.Errorf("expected: %v, act: %v", nil, err)
	}

	if id := user.GetID(); id != *mock.UserID {
		t.Errorf("expected: %v, act: %v", *mock.UserID, id)
	}

	user.ChangeName(mock.UserName)
	user.ChangePhone(mock.UserPhone)
	user.ChangePassport(mock.PassportEntity)
	user.ChangeDrivingLicense(mock.DrivingLicenseEntity)
	user.ChangeSnils(mock.UserSnils)
	user.ChangeInn(mock.UserInn)
	user.ChangePhoto(&mock.UserPhoto)
	user.ChangeCard(mock.UserCard)
	user.IncreaseRating()
	user.DecreaseRating()
	user.GrantAccess(*mock.UserOtherID, valuetypes.AccessorContacts)
	user.RevokeAccess(*mock.UserOtherID, valuetypes.AccessorContacts)
	user.GrantFullAccess(*mock.UserOtherID)
	user.RevokeFullAccess(*mock.UserOtherID)
	user.Block()
	user.Activate()
	user.Remove()

	if id := user.GetID(); id != *mock.UserID {
		t.Errorf("expected: %v, act: %v", *mock.UserID, id)
	}
}

func TestUserNil_String(t *testing.T) {
	expected := "User Aggregate ID 059b4e12-6983-4806-bd5a-cc3433e78f66\nRemoved: true\n"

	user, _ := NewUserNil(mock.UserID)

	if act := user.String(); act != expected {
		t.Errorf("expected: %v, act: %v", expected, act)
	}
}
