package valuetypes

import "testing"

func TestNewPassportID(t *testing.T) {
	var data = []struct {
		serial string
		number string
		err    error
	}{
		{"0000", "000000", nil},
		{"1234", "123456", nil},
		{"", "000000", ErrInvalidPassportSerial},
		{"000", "000000", ErrInvalidPassportSerial},
		{"0000", "", ErrInvalidPassportNumber},
		{"0000", "00000", ErrInvalidPassportNumber},
		{"", "", ErrInvalidPassportSerial},
		{"0", "0", ErrInvalidPassportSerial},
	}

	for _, val := range data {
		id, err := NewPassportID(val.serial, val.number)
		if err != val.err {
			t.Errorf("expected error: %v, act: %v", val.err, err)
		}

		if err != nil {
			continue
		}

		if id.serial != val.serial {
			t.Errorf("expected: %v, act: %v", val.serial, id.serial)
		}

		if id.number != val.number {
			t.Errorf("expected: %v, act: %v", val.number, id.number)
		}
	}
}

func TestPassportID_String(t *testing.T) {
	var data = []struct {
		serial   string
		number   string
		expected string
	}{
		{"0000", "000000", "00 00	000000"},
		{"1234", "123456", "12 34	123456"},
	}

	for _, val := range data {
		id, _ := NewPassportID(val.serial, val.number)
		if act := id.String(); act != val.expected {
			t.Errorf("expected: %v, act: %v", val.expected, act)
		}
	}
}