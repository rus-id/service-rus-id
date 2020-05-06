package valuetypes

const (
	A DrivingLicenseCategory = iota + 1
	B
	C
	D
)

type DrivingLicenseCategory int

func (d DrivingLicenseCategory) String() string {
	switch d {
	case A:
		return "A"
	case B:
		return "B"
	case C:
		return "C"
	case D:
		return "D"
	}

	return ""
}
