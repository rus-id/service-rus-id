package valuetypes

import (
	"testing"
)

func TestNewPhone(t *testing.T) {
	data := []struct {
		code   CountryCode
		number string
		err    error
	}{
		{RusCountryCode, "9039615321", nil},
		{RusCountryCode, "9267773311", nil},
		{9, "9267773311", ErrInvalidCountryCode},
		{RusCountryCode, "92677733", ErrInvalidPhoneNumber},
		{0, "", ErrInvalidCountryCode},
	}

	for _, val := range data {
		phone, err := NewPhone(val.code, val.number)

		if err != val.err {
			t.Errorf("expected error: %v, actual: %v", val.err, err)
		}

		if err != nil {
			continue
		}

		if phone.code != val.code {
			t.Errorf("expected: %v, actual: %v", val.code, phone.code)
		}

		if phone.number != val.number {
			t.Errorf("expected: %v, actual: %v", val.number, phone.number)
		}
	}
}

func TestPhone_Getters(t *testing.T) {
	data := []struct {
		code   CountryCode
		number string
	}{
		{RusCountryCode, "9039615321"},
		{RusCountryCode, "9267773311"},
	}

	for _, val := range data {
		phone, _ := NewPhone(val.code, val.number)

		if act := phone.GetCode(); act != val.code {
			t.Errorf("expected: %v, actual: %v", val.code, act)
		}

		if act := phone.GetNumber(); act != val.number {
			t.Errorf("expected: %v, actual: %v", val.number, act)
		}
	}
}

func TestPhone_String(t *testing.T) {
	exp := "+7(903)961-53-21"
	phone, _ := NewPhone(RusCountryCode, "9039615321")

	if act := phone.String(); act != exp {
		t.Errorf("expected: %v, actual: %v", exp, act)
	}
}
