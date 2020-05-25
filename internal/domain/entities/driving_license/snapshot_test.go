// +build entities

package driving_license_test

import (
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license"
)

func TestNewSnapshot(t *testing.T) {
	act := NewSnapshot(
		serial,
		number,
		int64(category),
		firstName,
		middleName,
		lastName,
		birthdayStamp,
		issuedStamp,
		expiresStamp,
		country,
		specialMarks,
		gibddValidation,
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

	if act.Birthday != birthdayStamp {
		t.Errorf("expected: %v, act: %v", birthdayStamp, act.Birthday)
	}

	if act.Issued != issuedStamp {
		t.Errorf("expected: %v, act: %v", issuedStamp, act.Issued)
	}

	if act.Expires != expiresStamp {
		t.Errorf("expected: %v, act: %v", expiresStamp, act.Expires)
	}

	if act.Residence != country {
		t.Errorf("expected: %v, act: %v", country, act.Residence)
	}

	if act.SpecialMarks != specialMarks {
		t.Errorf("expected: %v, act: %v", specialMarks, act.SpecialMarks)
	}

	if act.GibddValidation != gibddValidation {
		t.Errorf("expected: %v, act: %v", gibddValidation, act.GibddValidation)
	}

	if act.DocumentValidation != documentValidation {
		t.Errorf("expected: %v, act: %v", documentValidation, act.DocumentValidation)
	}
}

func TestLoadFromSnapshot_Success(t *testing.T) {
	snapshot := NewSnapshot(
		serial,
		number,
		int64(category),
		firstName,
		middleName,
		lastName,
		birthdayStamp,
		issuedStamp,
		expiresStamp,
		country,
		specialMarks,
		gibddValidation,
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

	if act := dl.GetCategory(); act != category {
		t.Errorf("expected: %v, act: %v", category, act)
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

	if act := dl.GetIssued(); act != issued {
		t.Errorf("expected: %v, act: %v", issued, act)
	}

	if act := dl.GetExpires(); act != expires {
		t.Errorf("expected: %v, act: %v", expires, act)
	}

	if act := dl.GetResidence().GetValue(); act != country {
		t.Errorf("expected: %v, act: %v", country, act)
	}

	if act := dl.GetSpecialMarks(); act != specialMarks {
		t.Errorf("expected: %v, act: %v", specialMarks, act)
	}

	if act := dl.GetValidation().GetGibdd(); act != gibddValidation {
		t.Errorf("expected: %v, act: %v", gibddValidation, act)
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
	dl, _ := NewDrivingLicense(
		id,
		category,
		name,
		&birthday,
		&issued,
		&expires,
		residence,
		specialMarks,
		validation)

	act := GetSnapshot(dl)

	if act.Serial != serial {
		t.Errorf("expected: %v, act: %v", serial, act.Serial)
	}

	if act.Number != number {
		t.Errorf("expected: %v, act: %v", number, act.Number)
	}

	if exp := int64(category); act.Category != exp {
		t.Errorf("expected: %v, act: %v", exp, act.Category)
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

	if act.Issued != issuedStamp {
		t.Errorf("expected: %v, act: %v", issuedStamp, act.Issued)
	}

	if act.Expires != expiresStamp {
		t.Errorf("expected: %v, act: %v", expiresStamp, act.Expires)
	}

	if act.Residence != country {
		t.Errorf("expected: %v, act: %v", country, act.Residence)
	}

	if act.SpecialMarks != specialMarks {
		t.Errorf("expected: %v, act: %v", specialMarks, act.SpecialMarks)
	}

	if act.GibddValidation != gibddValidation {
		t.Errorf("expected: %v, act: %v", gibddValidation, act.GibddValidation)
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
