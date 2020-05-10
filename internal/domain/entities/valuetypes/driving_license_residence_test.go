package valuetypes

import (
	"reflect"
	"testing"
)

func TestNewResidence(t *testing.T) {
	data := []struct {
		value    string
		expected *DrivingLicenseResidence
		err      error
	}{
		{"123", &DrivingLicenseResidence{"123"}, nil},
		{"", nil, ErrInvalidResidence},
	}

	for _, val := range data {
		act, err := NewResidence(val.value)
		if !reflect.DeepEqual(act, val.expected) {
			t.Errorf("expected: %v, act: %v", val.expected, act)
		}

		if err != val.err {
			t.Errorf("expected error: %v, act: %v", val.err, err)
		}
	}
}
