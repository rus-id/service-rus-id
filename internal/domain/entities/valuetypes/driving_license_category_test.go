package valuetypes

import "testing"

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
			t.Errorf("expected: %v, act: %v", val.expected, act)
		}
	}
}
