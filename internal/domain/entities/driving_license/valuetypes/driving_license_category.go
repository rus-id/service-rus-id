package valuetypes

const (
	DrivingLicenseA DrivingLicenseCategory = iota + 1
	DrivingLicenseB
	DrivingLicenseC
	DrivingLicenseD
)

type DrivingLicenseCategory int

func (d DrivingLicenseCategory) String() string {
	switch d {
	case DrivingLicenseA:
		return "A"
	case DrivingLicenseB:
		return "B"
	case DrivingLicenseC:
		return "C"
	case DrivingLicenseD:
		return "D"
	}

	return ""
}

func (d DrivingLicenseCategory) IsValid() bool {
	if d != DrivingLicenseA &&
		d != DrivingLicenseB &&
		d != DrivingLicenseC &&
		d != DrivingLicenseD {
		return false
	}

	return true
}
