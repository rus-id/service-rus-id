package valuetypes

import "testing"

var testDataInn = []struct {
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

func TestNewInn(t *testing.T) {
	for _, val := range testDataInn {
		inn, err := NewInn(val.inn)

		if val.expected && Inn(val.inn) != inn {
			t.Errorf("expected %v, actual %v. error: %v", val.inn, inn, err)
		}

		if !val.expected && err != ErrInvalidInn {
			t.Errorf("expected error %v, actual %v", ErrInvalidInn, err)
		}
	}
}

func TestValidateInn(t *testing.T) {
	for _, val := range testDataInn {
		ok, err := validateInn(val.inn)
		if ok != val.expected {
			t.Errorf("inn %v not valid. error: %v", val.inn, err)
		}
	}
}
