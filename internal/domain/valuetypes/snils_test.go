package valuetypes_test

import (
	"reflect"
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestNewSnils(t *testing.T) {
	var data = []struct {
		snils string
		err   error
	}{
		{"59650418527", nil},
		{"60319270458", nil},
		{"39650468527", ErrInvalidSnils},
		{"30319270422", ErrInvalidSnils},
	}

	for _, val := range data {
		snils, err := NewSnils(val.snils)

		if err != val.err {
			t.Errorf("expected error %v, actual: %v", val.err, err)
		}

		if err != nil {
			continue
		}

		if exp := Snils(val.snils); !reflect.DeepEqual(&exp, snils) {
			t.Errorf("expected: %v, actual: %v", exp, snils)
		}
	}
}

func TestValidateSnils(t *testing.T) {
	var data = []struct {
		snils    string
		expected bool
	}{
		{"59650418527", true},
		{"60319270458", true},
		{"39650468527", false},
		{"30319270422", false},
	}

	for _, val := range data {
		ok, err := ValidateSnils(val.snils)
		if ok != val.expected {
			t.Errorf("snils %v not valid. error: %v", val.snils, err)
		}
	}
}

func TestSnils_String(t *testing.T) {
	data := []struct {
		snils    string
		expected string
	}{
		{"59650418527", "596-504-185-27"},
		{"60319270458", "603-192-704-58"},
	}

	for _, val := range data {
		snils, _ := NewSnils(val.snils)

		if act := snils.String(); act != val.expected {
			t.Errorf("expected: %q, actual: %q", val.snils, snils)
		}
	}
}
