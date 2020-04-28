package valuetypes

import "testing"

var testDataSnils = []struct {
	snils    string
	expected bool
}{
	{"59650418527", true},
	{"60319270458", true},
	{"39650468527", false},
	{"30319270422", false},
}

func TestNewSnils(t *testing.T) {
	for _, val := range testDataSnils {
		snils, err := NewSnils(val.snils)

		if val.expected && Snils(val.snils) != snils {
			t.Errorf("expected %v, actual %v. error: %v", val.snils, snils, err)
		}

		if !val.expected && err != ErrInvalidSnils {
			t.Errorf("expected error %v, actual %v", ErrInvalidSnils, err)
		}
	}
}

func TestValidateSnils(t *testing.T) {
	for _, val := range testDataSnils {
		ok, err := validateSnils(val.snils)
		if ok != val.expected {
			t.Errorf("snils %v not valid. error: %v", val.snils, err)
		}
	}
}
