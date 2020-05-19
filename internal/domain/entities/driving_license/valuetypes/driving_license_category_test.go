package valuetypes_test

import (
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license/valuetypes"
)

func TestDrivingLicenseCategory_String(t *testing.T) {
	data := []struct {
		category DrivingLicenseCategory
		expected string
	}{
		{DrivingLicenseA, "A"},
		{DrivingLicenseB, "B"},
		{DrivingLicenseC, "C"},
		{DrivingLicenseD, "D"},
	}

	for _, val := range data {
		if act := val.category.String(); act != val.expected {
			t.Errorf("expected: %q, act: %q", val.expected, act)
		}
	}
}

func TestDrivingLicenseCategory_IsValid(t *testing.T) {
	data := []struct {
		category DrivingLicenseCategory
		isValid  bool
	}{
		{DrivingLicenseA, true},
		{DrivingLicenseB, true},
		{DrivingLicenseC, true},
		{DrivingLicenseD, true},
		{100, false},
		{-100, false},
	}

	for _, val := range data {
		if act := val.category.IsValid(); act != val.isValid {
			t.Errorf("expected: %v, act: %v", val.isValid, act)
		}
	}
}
