package valuetypes

import "fmt"

type DrivingLicenseValidation struct {
	gibdd    bool
	document bool
}

func NewDrivingLicenseValidation(gibdd, document bool) *DrivingLicenseValidation {
	return &DrivingLicenseValidation{gibdd: gibdd, document: document}
}

func (v DrivingLicenseValidation) GetGibdd() bool {
	return v.gibdd
}

func (v DrivingLicenseValidation) GetDocument() bool {
	return v.document
}

func (v DrivingLicenseValidation) String() string {
	isValid := func(valid bool) string {
		if valid {
			return "valid"
		}
		return "not valid"
	}

	return fmt.Sprintf("GIBDD %v; document %v;", isValid(v.gibdd), isValid(v.document))
}
