package snapshots

type DrivingLicenseSnapshot struct {
	Serial             string
	Number             string
	Category           int64
	FirstName          string
	MiddleName         *string
	LastName           string
	Birthday           int64
	Issued             int64
	Expired            int64
	Residence          string
	SpecialMarks       string
	GibddValidation    bool
	DocumentValidation bool
	Timestamp          int64
}

func NewDrivingLicense(
	serial string,
	number string,
	category int64,
	firstName string,
	middleName *string,
	lastName string,
	birthday int64,
	issued int64,
	expired int64,
	residence string,
	specialMarks string,
	gibddValidation bool,
	documentValidation bool,
	timestamp int64,
) DrivingLicenseSnapshot {
	return DrivingLicenseSnapshot{
		Serial:             serial,
		Number:             number,
		Category:           category,
		FirstName:          firstName,
		MiddleName:         middleName,
		LastName:           lastName,
		Birthday:           birthday,
		Issued:             issued,
		Expired:            expired,
		Residence:          residence,
		SpecialMarks:       specialMarks,
		GibddValidation:    gibddValidation,
		DocumentValidation: documentValidation,
		Timestamp:          timestamp,
	}
}
