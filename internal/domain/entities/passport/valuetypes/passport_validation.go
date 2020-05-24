package valuetypes

import "fmt"

type PassportValidation struct {
	ufms     bool
	mvd      bool
	fssp     bool
	document bool
}

func NewPassportValidation(ufms, mvd, fssp, document bool) *PassportValidation {
	return &PassportValidation{ufms: ufms, mvd: mvd, fssp: fssp, document: document}
}

func (v PassportValidation) GetUfms() bool {
	return v.ufms
}

func (v PassportValidation) GetMvd() bool {
	return v.mvd
}

func (v PassportValidation) GetFssp() bool {
	return v.fssp
}

func (v PassportValidation) GetDocument() bool {
	return v.document
}

func (v PassportValidation) String() string {
	return fmt.Sprintf("UFMS valid: %v. MVD valid: %v. FSSP valid: %v Document valid: %v",
		v.ufms,
		v.mvd,
		v.fssp,
		v.document)
}
