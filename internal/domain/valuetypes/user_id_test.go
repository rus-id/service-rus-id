package valuetypes_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestNewUserID_Success(t *testing.T) {
	rawUserID, _ := uuid.Parse("059b4e12-6983-4806-bd5a-cc3433e78f66")

	act, err := NewUserID(rawUserID)

	if err != nil {
		t.Errorf("error occured: %v", err)
	}

	if reflect.DeepEqual(act, rawUserID) {
		t.Errorf("expected: %v, act: %v", rawUserID, act)
	}
}

func TestCreateNewUserID(t *testing.T) {
	emptyID := fmt.Sprint([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	id := CreateUserID()

	if act := fmt.Sprint(id); act == emptyID {
		t.Errorf("empty identity: %v", act)
	}
}

func TestUserID_String(t *testing.T) {
	const expected = "059b4e12-6983-4806-bd5a-cc3433e78f66"
	rawUserID, _ := uuid.Parse(expected)
	act, _ := NewUserID(rawUserID)

	if act.String() != expected {
		t.Errorf("expected: %v, act: %v", expected, act)
	}
}
