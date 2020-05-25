package passport

import (
	"time"

	passValueTypes "github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

type Snapshot struct {
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

func NewSnapshot(
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
) Snapshot {
	return Snapshot{
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
	}
}

func GetSnapshot(passport *Passport) *Snapshot {
	if passport == nil {
		return nil
	}

	snapshot := NewSnapshot(
		passport.GetID().GetSerial(),
		passport.GetID().GetNumber(),
		passport.GetName().GetFirst(),
		passport.GetName().GetMiddle(),
		passport.GetName().GetLast(),
		passport.GetBirthday().Unix(),
		passport.GetIssued().GetOrganisation(),
		passport.GetIssued().GetDate().Unix(),
		passport.GetIssued().GetCode(),
		string(passport.GetRegistration()),
		passport.GetValidation().GetUfms(),
		passport.GetValidation().GetMvd(),
		passport.GetValidation().GetFssp(),
		passport.GetValidation().GetDocument(),
	)

	return &snapshot
}

func LoadFromSnapshot(snapshot *Snapshot) (*Passport, error) {
	if snapshot == nil {
		return nil, nil
	}

	id, err := passValueTypes.NewPassportID(snapshot.Serial, snapshot.Number)
	if err != nil {
		return nil, err
	}

	name, err := valuetypes.NewName(snapshot.FirstName, snapshot.MiddleName, snapshot.LastName)
	if err != nil {
		return nil, err
	}

	birthday := time.Unix(snapshot.Birthday, 0).UTC()
	issueDate := time.Unix(snapshot.IssuedDate, 0).UTC()
	issue, err := passValueTypes.NewPassportIssue(snapshot.IssuedOrganisation, issueDate, snapshot.IssuedCode)
	if err != nil {
		return nil, err
	}

	registration, err := valuetypes.NewAddress(snapshot.Registration)
	if err != nil {
		return nil, err
	}

	validation := passValueTypes.NewPassportValidation(
		snapshot.UfmsValidation,
		snapshot.MvdValidation,
		snapshot.FsspValidation,
		snapshot.DocumentValidation)

	return NewPassport(id, name, &birthday, issue, registration, validation)
}
