package valuetypes

import "fmt"

type DrivingLicenseValidation struct {
	gibdd    bool
	document bool
}

func NewDrivingLicenseValidation(gibdd, document bool) *DrivingLicenseValidation {
	return &DrivingLicenseValidation{gibdd: gibdd, document: document}
}

func (v *DrivingLicenseValidation) GetGibdd() bool {
	return v.gibdd
}

func (v *DrivingLicenseValidation) GetDocument() bool {
	return v.document
}

func (v *DrivingLicenseValidation) String() string {
	return fmt.Sprintf("GIBDD valid: %v. Document valid: %v", v.gibdd, v.document)
}
