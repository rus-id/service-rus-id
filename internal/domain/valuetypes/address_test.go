package valuetypes_test

import (
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestNewAddress(t *testing.T) {
	data := []struct {
		address  string
		expected Address
		err      error
	}{
		{"Москва, Шипиловская улица", "Москва, Шипиловская улица", nil},
		{"123", "123", nil},
		{"", "", ErrInvalidAddress},
	}

	for _, val := range data {
		act, err := NewAddress(val.address)
		if act != val.expected {
			t.Errorf("expected: %v, act: %v", val.expected, act)
		}

		if err != val.err {
			t.Errorf("expected error: %v, act: %v", val.err, err)
		}
	}
}
