package valuetypes_test

import (
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestUserStatus_String(t *testing.T) {
	data := []struct {
		state    UserState
		expected string
	}{
		{UserStateActive, "active user"},
		{UserStateBlocked, "blocked user"},
		{777, ""},
	}

	for _, val := range data {
		if act := val.state.String(); act != val.expected {
			t.Errorf("expected: %q, act: %q", val.expected, act)
		}
	}
}

func TestUserStatus_IsValid(t *testing.T) {
	data := []struct {
		state   UserState
		isValid bool
	}{
		{UserStateActive, true},
		{UserStateBlocked, true},
		{100, false},
		{-100, false},
	}

	for _, val := range data {
		if act := val.state.IsValid(); act != val.isValid {
			t.Errorf("expected: %v, act: %v", val.isValid, act)
		}
	}
}
