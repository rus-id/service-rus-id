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
	isValid := func(valid bool) string {
		if valid {
			return "valid"
		}
		return "not valid"
	}

	return fmt.Sprintf("UFMS %v; MVD %v; FSSP %v; document %v;",
		isValid(v.ufms),
		isValid(v.mvd),
		isValid(v.fssp),
		isValid(v.document))
}
