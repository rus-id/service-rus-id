package snapshots

type PassportSnapshot struct {
	Serial             string
	Number             string
	FirstName          string
	MiddleName         *string
	LastName           string
	Birthday           int64
	IssuedOrganisation string
	IssuedDate         int64
	IssuedCode         string
	Registration       string
	UfmsValidation     bool
	MvdValidation      bool
	FsspValidation     bool
	DocumentValidation bool
	Timestamp          int64
}

func NewPassport(
	serial string,
	number string,
	firstName string,
	middleName *string,
	lastName string,
	birthday int64,
	issuedOrganisation string,
	issuedDate int64,
	issuedCode string,
	registration string,
	ufmsValidation bool,
	mvdValidation bool,
	fsspValidation bool,
	documentValidation bool,
	timestamp int64,
) PassportSnapshot {
	return PassportSnapshot{
		Serial:             serial,
		Number:             number,
		FirstName:          firstName,
		MiddleName:         middleName,
		LastName:           lastName,
		Birthday:           birthday,
		IssuedOrganisation: issuedOrganisation,
		IssuedDate:         issuedDate,
		IssuedCode:         issuedCode,
		Registration:       registration,
		UfmsValidation:     ufmsValidation,
		MvdValidation:      mvdValidation,
		FsspValidation:     fsspValidation,
		DocumentValidation: documentValidation,
		Timestamp:          timestamp,
	}
}
