package valuetypes_test

import (
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestNewName_Success(t *testing.T) {
	middle := "B"
	emptyMiddle := ""

	data := []struct {
		first  string
		middle string
		last   string
	}{
		{"Boris", middle, "Goldovsky"},
		{"Boris", emptyMiddle, "Goldovsky"},
		{"Boris", "", "Goldovsky"},
	}

	for _, val := range data {
		name, err := NewName(val.first, &val.middle, val.last)

		if err != nil {
			t.Errorf("error occured: %v", err)
		}

		if act := name.GetFirst(); act != val.first {
			t.Errorf("expected: %v, act: %v", val.first, act)
		}

		if act := name.GetMiddle(); act != val.middle {
			t.Errorf("expected: %v, act: %v", val.middle, act)
		}

		if act := name.GetLast(); act != val.last {
			t.Errorf("expected: %v, act: %v", val.last, act)
		}
	}
}

func TestNewName_Error(t *testing.T) {
	data := []struct {
		first    string
		last     string
		expected error
	}{
		{"", "Goldovsky", ErrInvalidFirstName},
		{"Boris", "", ErrInvalidLastName},
		{"", "", ErrInvalidFirstName},
	}

	for _, val := range data {
		_, err := NewName(val.first, nil, val.last)

		if err != val.expected {
			t.Errorf("expected error: %v, act: %v", val.expected, err)
		}
	}
}

func TestName_Getters(t *testing.T) {
	middle := "B"
	emptyMiddle := ""

	data := []struct {
		first          string
		middle         *string
		expectedMiddle string
		last           string
	}{
		{"Boris", &middle, middle, "Goldovsky"},
		{"Boris", &emptyMiddle, emptyMiddle, "Goldovsky"},
		{"Boris", nil, "", "Goldovsky"},
	}

	for _, val := range data {
		name, _ := NewName(val.first, val.middle, val.last)

		if act := name.GetFirst(); act != val.first {
			t.Errorf("expected: %v, act: %v", val.last, act)
		}

		if act := name.GetMiddle(); act != val.expectedMiddle {
			t.Errorf("expected: %v, act: %v", val.middle, act)
		}

		if act := name.GetLast(); act != val.last {
			t.Errorf("expected: %v, act: %v", val.last, act)
		}
	}
}

func TestName_String(t *testing.T) {
	middle := "B"
	emptyMiddle := ""

	data := []struct {
		first    string
		middle   *string
		last     string
		expected string
	}{
		{"Boris", &middle, "Goldovsky", "Boris B Goldovsky"},
		{"Boris", &emptyMiddle, "Goldovsky", "Boris Goldovsky"},
		{"Boris", nil, "Goldovsky", "Boris Goldovsky"},
	}

	for _, val := range data {
		name, _ := NewName(val.first, val.middle, val.last)

		act := name.String()
		if act != val.expected {
			t.Errorf("expected: %q, act: %q", val.expected, act)
		}
	}
}
