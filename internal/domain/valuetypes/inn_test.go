package valuetypes_test

import (
	"reflect"
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestNewInn(t *testing.T) {
	var data = []struct {
		inn string
		err error
	}{
		{"926902267890", nil},
		{"889373498613", nil},
		{"326902237890", ErrInvalidInn},
		{"389373498413", ErrInvalidInn},

		{"8204347950", nil},
		{"1131784600", nil},
		{"3444347950", ErrInvalidInn},
		{"3317846540", ErrInvalidInn},
	}

	for _, val := range data {
		inn, err := NewInn(val.inn)

		if err != val.err {
			t.Errorf("expected error %v, actual %v", val.err, err)
		}

		if err != nil {
			continue
		}

		if exp := Inn(val.inn); !reflect.DeepEqual(&exp, inn) {
			t.Errorf("expected %v, actual %v", exp, inn)
		}
	}
}

func TestValidateInn(t *testing.T) {
	var data = []struct {
		inn      string
		expected bool
	}{
		{"926902267890", true},
		{"889373498613", true},
		{"326902237890", false},
		{"389373498413", false},

		{"8204347950", true},
		{"1131784600", true},
		{"3444347950", false},
		{"3317846540", false},
	}

	for _, val := range data {
		ok, err := ValidateInn(val.inn)
		if ok && err != nil {
			t.Errorf("inn %v not valid. error: %v", val.inn, err)
		}
	}
}

func TestInn_String(t *testing.T) {
	data := []string{"926902267890", "889373498613"}

	for _, val := range data {
		inn, _ := NewInn(val)
		if act := inn.String(); act != val {
			t.Errorf("expected: %q, actual: %q", val, inn)
		}
	}
}
