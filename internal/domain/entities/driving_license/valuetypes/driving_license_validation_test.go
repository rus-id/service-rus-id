package valuetypes_test

import (
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license/valuetypes"
)

func TestDrivingLicenseValidation_Getters(t *testing.T) {
	data := []struct {
		gibdd    bool
		document bool
	}{
		{false, false},
		{true, false},
		{false, true},
		{true, true},
	}

	for _, val := range data {
		validation := NewDrivingLicenseValidation(val.gibdd, val.document)

		if act := validation.GetGibdd(); act != val.gibdd {
			t.Errorf("expected gibdd: %v, act: %v", val.gibdd, act)
		}

		if act := validation.GetDocument(); act != val.document {
			t.Errorf("expected document: %v, act: %v", val.document, act)
		}
	}
}

func TestDrivingLicenseValidation_String(t *testing.T) {
	const expected = "GIBDD valid; document not valid;"

	act := NewDrivingLicenseValidation(true, false).String()

	if act != expected {
		t.Errorf("expected text: %q, act: %q", expected, act)
	}
}
