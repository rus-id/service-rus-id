package driving_license

import (
	"time"

	dlValueTypes "github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

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
}

func NewSnapshot(
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
	}
}

func GetSnapshot(drivingLicense *DrivingLicense) *DrivingLicenseSnapshot {
	if drivingLicense == nil {
		return nil
	}

	snapshot := NewSnapshot(
		drivingLicense.GetID().GetSerial(),
		drivingLicense.GetID().GetNumber(),
		int64(drivingLicense.GetCategory()),
		drivingLicense.GetName().GetFirst(),
		drivingLicense.GetName().GetMiddle(),
		drivingLicense.GetName().GetLast(),
		drivingLicense.GetBirthday().Unix(),
		drivingLicense.GetIssued().Unix(),
		drivingLicense.GetExpired().Unix(),
		drivingLicense.GetResidence().GetValue(),
		drivingLicense.GetSpecialMarks(),
		drivingLicense.GetValidation().GetGibdd(),
		drivingLicense.GetValidation().GetDocument(),
	)

	return &snapshot
}

func LoadFromSnapshot(snapshot *DrivingLicenseSnapshot) (*DrivingLicense, error) {
	if snapshot == nil {
		return nil, nil
	}

	id, err := dlValueTypes.NewDrivingLicenseID(snapshot.Serial, snapshot.Number)
	if err != nil {
		return nil, err
	}

	name, err := valuetypes.NewName(snapshot.FirstName, snapshot.MiddleName, snapshot.LastName)
	if err != nil {
		return nil, err
	}

	category := dlValueTypes.DrivingLicenseCategory(snapshot.Category)
	birthday := time.Unix(snapshot.Birthday, 0)
	issued := time.Unix(snapshot.Issued, 0)
	expired := time.Unix(snapshot.Expired, 0)

	residence, err := dlValueTypes.NewResidence(snapshot.Residence)
	if err != nil {
		return nil, err
	}

	validation := dlValueTypes.NewDrivingLicenseValidation(snapshot.GibddValidation, snapshot.DocumentValidation)

	return NewDrivingLicense(
		id,
		category,
		name,
		&birthday,
		&issued,
		&expired,
		residence,
		snapshot.SpecialMarks,
		validation)
}
