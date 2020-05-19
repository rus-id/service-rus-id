// +build entities

package driving_license_test

import (
	"testing"
	"time"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license/valuetypes"
	commonTypes "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

var (
	id, _        = valuetypes.NewDrivingLicenseID("7777", "777777")
	category     = valuetypes.DrivingLicenseA
	name, _      = commonTypes.NewName("Boris", nil, "Goldovsky")
	birthday     = time.Date(1986, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	issued       = time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	expired      = time.Date(2025, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	residence, _ = valuetypes.NewResidence("Russia")
	specialMarks = "empty mark"
	validation   = valuetypes.NewDrivingLicenseValidation(true, false)
)

func TestNewDrivingLicense_Success(t *testing.T) {
	dl, err := NewDrivingLicense(
		id,
		category,
		name,
		&birthday,
		&issued,
		&expired,
		residence,
		specialMarks,
		validation)

	if dl == nil {
		t.Errorf("driver license must be not null")
	}

	if err != nil {
		t.Errorf("driver license error must be null")
	}
}

func TestNewDrivingLicense_Errors(t *testing.T) {
	data := []struct {
		id         *valuetypes.DrivingLicenseID
		category   valuetypes.DrivingLicenseCategory
		name       *commonTypes.Name
		birthday   *time.Time
		issued     *time.Time
		expired    *time.Time
		residence  *valuetypes.DrivingLicenseResidence
		validation *valuetypes.DrivingLicenseValidation
		err        error
	}{
		{
			id,
			category,
			name,
			&birthday,
			&issued,
			&expired,
			residence,
			validation,
			nil,
		},
		{
			nil,
			category,
			name,
			&birthday,
			&issued,
			&expired,
			residence,
			validation,
			ErrInvalidID,
		},
		{
			id,
			123,
			name,
			&birthday,
			&issued,
			&expired,
			residence,
			validation,
			ErrInvalidCategory,
		},
		{
			id,
			category,
			nil,
			&birthday,
			&issued,
			&expired,
			residence,
			validation,
			ErrInvalidName,
		},
		{
			id,
			category,
			name,
			nil,
			&issued,
			&expired,
			residence,
			validation,
			ErrInvalidBirthday,
		},
		{
			id,
			category,
			name,
			&birthday,
			nil,
			&expired,
			residence,
			validation,
			ErrInvalidIssue,
		},
		{
			id,
			category,
			name,
			&birthday,
			&issued,
			&expired,
			nil,
			validation,
			ErrInvalidResidence,
		},
		{
			id,
			category,
			name,
			&birthday,
			&issued,
			&expired,
			residence,
			nil,
			ErrInvalidValidation,
		},
		{
			id,
			category,
			name,
			&birthday,
			&expired,
			&issued,
			residence,
			validation,
			ErrInvalidDates,
		},
	}

	for _, val := range data {
		_, err := NewDrivingLicense(
			val.id,
			val.category,
			val.name,
			val.birthday,
			val.issued,
			val.expired,
			val.residence,
			specialMarks,
			val.validation)

		if err != val.err {
			t.Errorf("expected err: %v, act: %v", val.err, err)
		}
	}
}

func TestDrivingLicense_Getters(t *testing.T) {
	dl, _ := NewDrivingLicense(
		id,
		category,
		name,
		&birthday,
		&issued,
		&expired,
		residence,
		specialMarks,
		validation)

	if act := dl.GetID(); act != *id {
		t.Errorf("expected: %v, act: %v", id, act)
	}

	if act := dl.GetCategory(); act != category {
		t.Errorf("expected: %v, act: %v", category, act)
	}

	if act := dl.GetName(); act != *name {
		t.Errorf("expected: %v, act: %v", name, act)
	}

	if act := dl.GetBirthday(); act != birthday {
		t.Errorf("expected: %v, act: %v", &birthday, act)
	}

	if act := dl.GetIssued(); act != issued {
		t.Errorf("expected: %v, act: %v", &issued, act)
	}

	if act := dl.GetExpired(); act != expired {
		t.Errorf("expected: %v, act: %v", &expired, act)
	}

	if act := dl.GetResidence(); act != *residence {
		t.Errorf("expected: %v, act: %v", residence, act)
	}

	if act := dl.GetSpecialMarks(); act != specialMarks {
		t.Errorf("expected: %v, act: %v", specialMarks, act)
	}

	if act := dl.GetValidation(); act != *validation {
		t.Errorf("expected: %v, act: %v", validation, act)
	}
}

func TestDrivingLicense_Setters(t *testing.T) {
	newCategory := valuetypes.DrivingLicenseB
	newName, _ := commonTypes.NewName("Edward", nil, "Kondratev")
	newBirthday := time.Date(1988, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	newIssued := time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	newExpired := time.Date(2027, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	newResidence, _ := valuetypes.NewResidence("USA")
	newSpecialMarks := "new special mark"
	newValidation := valuetypes.NewDrivingLicenseValidation(false, true)

	dl, _ := NewDrivingLicense(
		id,
		category,
		name,
		&birthday,
		&issued,
		&expired,
		residence,
		specialMarks,
		validation)

	dl.ChangeCategory(newCategory)
	if act := dl.GetCategory(); act != newCategory {
		t.Errorf("expected: %v, act: %v", newCategory, act)
	}

	dl.ChangeName(newName)
	if act := dl.GetName(); act != *newName {
		t.Errorf("expected: %v, act: %v", newName, act)
	}

	dl.ChangeBirthday(&newBirthday)
	if act := dl.GetBirthday(); act != newBirthday {
		t.Errorf("expected: %v, act: %v", &newBirthday, act)
	}

	dl.ChangeIssued(&newIssued)
	if act := dl.GetIssued(); act != newIssued {
		t.Errorf("expected: %v, act: %v", &newIssued, act)
	}

	dl.ChangeExpired(&newExpired)
	if act := dl.GetExpired(); act != newExpired {
		t.Errorf("expected: %v, act: %v", &newExpired, act)
	}

	dl.ChangeResidence(newResidence)
	if act := dl.GetResidence(); act != *newResidence {
		t.Errorf("expected: %v, act: %v", newResidence, act)
	}

	dl.ChangeSpecialMark(newSpecialMarks)
	if act := dl.GetSpecialMarks(); act != newSpecialMarks {
		t.Errorf("expected: %v, act: %v", newSpecialMarks, act)
	}

	dl.ChangeValidation(newValidation)
	if act := dl.GetValidation(); act != *newValidation {
		t.Errorf("expected: %v, act: %v", newValidation, act)
	}
}
