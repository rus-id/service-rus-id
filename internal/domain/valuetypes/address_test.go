package valuetypes_test

import (
	"reflect"
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestNewAddress(t *testing.T) {
	addr1 := Address("Москва, Шипиловская улица")
	addr2 := Address("123")

	data := []struct {
		address  string
		expected *Address
		err      error
	}{
		{"Москва, Шипиловская улица", &addr1, nil},
		{"123", &addr2, nil},
		{"", nil, ErrInvalidAddress},
	}

	for _, val := range data {
		act, err := NewAddress(val.address)
		if err != val.err {
			t.Errorf("expected error: %v, act: %v", val.err, err)
		}

		if !reflect.DeepEqual(act, val.expected) {
			t.Errorf("expected: %v, act: %v", val.expected, act)
		}
	}
}
