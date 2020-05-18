package valuetypes_test

import (
	"reflect"
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license/valuetypes"
)

func TestNewResidence(t *testing.T) {
	residence, _ := NewResidence("123")

	data := []struct {
		value    string
		expected *DrivingLicenseResidence
		err      error
	}{
		{"123", residence, nil},
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
