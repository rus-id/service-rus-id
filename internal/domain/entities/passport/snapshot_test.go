package passport_test

import (
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"
)

func TestNewSnapshot(t *testing.T) {
	act := NewSnapshot(
		serial,
		number,
		firstName,
		middleName,
		lastName,
		birthdayStamp,
		issuedOrganisation,
		issuedStamp,
		issuedCode,
		address,
		ufmsValidation,
		mvdValidation,
		fsspValidation,
		documentValidation)

	if act.Serial != serial {
		t.Errorf("expected: %v, act: %v", serial, act.Serial)
	}

	if act.Number != number {
		t.Errorf("expected: %v, act: %v", number, act.Number)
	}

	if act.FirstName != firstName {
		t.Errorf("expected: %v, act: %v", firstName, act.FirstName)
	}

	if act.MiddleName != middleName {
		t.Errorf("expected: %v, act: %v", middleName, act.MiddleName)
	}

	if act.LastName != lastName {
		t.Errorf("expected: %v, act: %v", lastName, act.LastName)
	}

	if act.IssuedOrganisation != issuedOrganisation {
		t.Errorf("expected: %v, act: %v", issuedOrganisation, act.IssuedOrganisation)
	}

	if act.IssuedDate != issuedStamp {
		t.Errorf("expected: %v, act: %v", issuedStamp, act.IssuedDate)
	}

	if act.IssuedCode != issuedCode {
		t.Errorf("expected: %v, act: %v", issuedCode, act.IssuedCode)
	}

	if act.Registration != address {
		t.Errorf("expected: %v, act: %v", address, act.Registration)
	}

	if act.UfmsValidation != ufmsValidation {
		t.Errorf("expected: %v, act: %v", ufmsValidation, act.UfmsValidation)
	}

	if act.MvdValidation != mvdValidation {
		t.Errorf("expected: %v, act: %v", mvdValidation, act.MvdValidation)
	}

	if act.FsspValidation != fsspValidation {
		t.Errorf("expected: %v, act: %v", fsspValidation, act.FsspValidation)
	}

	if act.DocumentValidation != documentValidation {
		t.Errorf("expected: %v, act: %v", documentValidation, act.DocumentValidation)
	}
}

func TestLoadFromSnapshot_Success(t *testing.T) {
	snapshot := NewSnapshot(
		serial,
		number,
		firstName,
		middleName,
		lastName,
		birthdayStamp,
		issuedOrganisation,
		issuedStamp,
		issuedCode,
		address,
		ufmsValidation,
		mvdValidation,
		fsspValidation,
		documentValidation)

	dl, err := LoadFromSnapshot(&snapshot)
	if err != nil {
		t.Errorf("expected not error, act %v", err)
	}

	if act := dl.GetID().GetSerial(); act != serial {
		t.Errorf("expected: %v, act: %v", serial, act)
	}

	if act := dl.GetID().GetNumber(); act != number {
		t.Errorf("expected: %v, act: %v", number, act)
	}

	if act := dl.GetName().GetFirst(); act != firstName {
		t.Errorf("expected: %v, act: %v", firstName, act)
	}

	if act := dl.GetName().GetMiddle(); act != middleName {
		t.Errorf("expected: %v, act: %v", middleName, act)
	}

	if act := dl.GetName().GetLast(); act != lastName {
		t.Errorf("expected: %v, act: %v", lastName, act)
	}

	if act := dl.GetBirthday(); act != birthday {
		t.Errorf("expected: %v, act: %v", birthday, act)
	}

	if act := dl.GetIssued().GetOrganisation(); act != issuedOrganisation {
		t.Errorf("expected: %v, act: %v", issuedOrganisation, act)
	}

	if act := dl.GetIssued().GetDate(); act != issuedDate {
		t.Errorf("expected: %v, act: %v", issuedDate, act)
	}

	if act := dl.GetIssued().GetCode(); act != issuedCode {
		t.Errorf("expected: %v, act: %v", issuedCode, act)
	}

	if act := dl.GetRegistration(); act != *registration {
		t.Errorf("expected: %v, act: %v", *registration, act)
	}

	if act := dl.GetValidation().GetUfms(); act != ufmsValidation {
		t.Errorf("expected: %v, act: %v", ufmsValidation, act)
	}

	if act := dl.GetValidation().GetMvd(); act != mvdValidation {
		t.Errorf("expected: %v, act: %v", mvdValidation, act)
	}

	if act := dl.GetValidation().GetFssp(); act != fsspValidation {
		t.Errorf("expected: %v, act: %v", fsspValidation, act)
	}

	if act := dl.GetValidation().GetDocument(); act != documentValidation {
		t.Errorf("expected: %v, act: %v", documentValidation, act)
	}
}

func TestLoadFromSnapshot_Nil(t *testing.T) {
	act, err := LoadFromSnapshot(nil)
	if err != nil {
		t.Errorf("expected not error, act %v", err)
	}

	if act != nil {
		t.Errorf("expected: %v, act: %v", nil, act)
	}
}

func TestGetSnapshot_Success(t *testing.T) {
	passport, _ := NewPassport(
		id,
		name,
		&birthday,
		issued,
		registration,
		validation)

	act := GetSnapshot(passport)

	if act.Serial != serial {
		t.Errorf("expected: %v, act: %v", serial, act.Serial)
	}

	if act.Number != number {
		t.Errorf("expected: %v, act: %v", number, act.Number)
	}

	if act.FirstName != firstName {
		t.Errorf("expected: %v, act: %v", firstName, act.FirstName)
	}

	if act.MiddleName != middleName {
		t.Errorf("expected: %v, act: %v", middleName, act.MiddleName)
	}

	if act.LastName != lastName {
		t.Errorf("expected: %v, act: %v", lastName, act.LastName)
	}

	if act.Birthday != birthdayStamp {
		t.Errorf("expected: %v, act: %v", birthdayStamp, act.Birthday)
	}

	if act.IssuedOrganisation != issuedOrganisation {
		t.Errorf("expected: %v, act: %v", issuedOrganisation, act.IssuedOrganisation)
	}

	if act.IssuedDate != issuedStamp {
		t.Errorf("expected: %v, act: %v", issuedStamp, act.IssuedDate)
	}

	if act.IssuedCode != issuedCode {
		t.Errorf("expected: %v, act: %v", issuedCode, act.IssuedCode)
	}

	if act.Registration != address {
		t.Errorf("expected: %v, act: %v", address, act.Registration)
	}

	if act.UfmsValidation != ufmsValidation {
		t.Errorf("expected: %v, act: %v", ufmsValidation, act.UfmsValidation)
	}

	if act.MvdValidation != mvdValidation {
		t.Errorf("expected: %v, act: %v", mvdValidation, act.MvdValidation)
	}

	if act.FsspValidation != fsspValidation {
		t.Errorf("expected: %v, act: %v", fsspValidation, act.FsspValidation)
	}

	if act.DocumentValidation != documentValidation {
		t.Errorf("expected: %v, act: %v", documentValidation, act.DocumentValidation)
	}
}

func TestGetSnapshot_Nil(t *testing.T) {
	snapshot := GetSnapshot(nil)

	if snapshot != nil {
		t.Errorf("expected: %v, act: %v", nil, snapshot.Serial)
	}
}
