package valuetypes_test

import (
	"fmt"
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestNewUserID_Success(t *testing.T) {
	expected := fmt.Sprint([]byte{5, 155, 78, 18, 105, 131, 72, 6, 189, 90, 204, 52, 51, 231, 143, 102})

	id, err := NewUserID("059b4e12-6983-4806-bd5a-cc3433e78f66")

	if err != nil {
		t.Errorf("error occured: %v", err)
	}

	if act := fmt.Sprint(id); act != expected {
		t.Errorf("expected: %v, act: %v", expected, act)
	}
}

func TestNewUserID_Fail(t *testing.T) {
	expected := fmt.Sprint([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	id, err := NewUserID("QWERTY")

	if err == nil {
		t.Error("error expected, act nil")
	}

	if act := fmt.Sprint(id); act != expected {
		t.Errorf("expected: %v, act: %v", expected, act)
	}
}

func TestCreateNewUserID(t *testing.T) {
	emptyID := fmt.Sprint([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	id := CreateUserID()

	if act := fmt.Sprint(id); act == emptyID {
		t.Errorf("empty identity: %v", act)
	}
}
