// +build entities

package passport_test

import (
	"testing"
	"time"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport/valuetypes"
	commonTypes "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestNewPassport_Success(t *testing.T) {
	passport, err := NewPassport(
		id,
		name,
		&birthday,
		issued,
		registration,
		validation)

	if passport == nil {
		t.Errorf("passport must be not null")
	}

	if err != nil {
		t.Errorf("passport error must be null")
	}
}

func TestNewPassport_Errors(t *testing.T) {
	data := []struct {
		id           *valuetypes.PassportID
		name         *commonTypes.Name
		birthday     *time.Time
		issued       *valuetypes.PassportIssue
		registration *commonTypes.Address
		validation   *valuetypes.PassportValidation
		err          error
	}{
		{
			id,
			name,
			&birthday,
			issued,
			registration,
			validation,
			nil,
		},
		{
			nil,
			name,
			&birthday,
			issued,
			registration,
			validation,
			ErrInvalidID,
		},
		{
			id,
			nil,
			&birthday,
			issued,
			registration,
			validation,
			ErrInvalidName,
		},
		{
			id,
			name,
			nil,
			issued,
			registration,
			validation,
			ErrInvalidBirthday,
		},
		{
			id,
			name,
			&birthday,
			nil,
			registration,
			validation,
			ErrInvalidIssue,
		},
		{
			id,
			name,
			&birthday,
			issued,
			nil,
			validation,
			ErrInvalidRegistration,
		},
		{
			id,
			name,
			&birthday,
			issued,
			registration,
			nil,
			ErrInvalidValidation,
		},
	}

	for _, val := range data {
		_, err := NewPassport(
			val.id,
			val.name,
			val.birthday,
			val.issued,
			val.registration,
			val.validation)

		if err != val.err {
			t.Errorf("expected err: %v, act: %v", val.err, err)
		}
	}
}

func TestPassport_Getters(t *testing.T) {
	passport, _ := NewPassport(
		id,
		name,
		&birthday,
		issued,
		registration,
		validation)

	if act := passport.GetID(); act != *id {
		t.Errorf("expected: %v, act: %v", id, act)
	}

	if act := passport.GetName(); act != *name {
		t.Errorf("expected: %v, act: %v", name, act)
	}

	if act := passport.GetBirthday(); act != birthday {
		t.Errorf("expected: %v, act: %v", &birthday, act)
	}

	if act := passport.GetIssued(); act != *issued {
		t.Errorf("expected: %v, act: %v", &issued, act)
	}

	if act := passport.GetRegistration(); act != *registration {
		t.Errorf("expected: %v, act: %v", registration, act)
	}

	if act := passport.GetValidation(); act != *validation {
		t.Errorf("expected: %v, act: %v", validation, act)
	}
}

func TestPassport_Setters(t *testing.T) {
	newName, _ := commonTypes.NewName("Boris", nil, "Goldovsky")
	newBirthday := time.Date(1986, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	newIssued, _ := valuetypes.NewPassportIssue(
		"MVD",
		time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC),
		"770-77")
	newRegistration, _ := commonTypes.NewAddress("Russia, Moscow")
	newValidation := valuetypes.NewPassportValidation(true, false, true, false)

	passport, _ := NewPassport(
		id,
		name,
		&birthday,
		issued,
		registration,
		validation)

	passport.ChangeName(newName)
	if act := passport.GetName(); act != *newName {
		t.Errorf("expected: %v, act: %v", newName, act)
	}

	passport.ChangeBirthday(&newBirthday)
	if act := passport.GetBirthday(); act != newBirthday {
		t.Errorf("expected: %v, act: %v", &newBirthday, act)
	}

	passport.ChangeIssued(newIssued)
	if act := passport.GetIssued(); act != *newIssued {
		t.Errorf("expected: %v, act: %v", &newIssued, act)
	}

	passport.ChangeRegistration(newRegistration)
	if act := passport.GetRegistration(); act != *newRegistration {
		t.Errorf("expected: %v, act: %v", newRegistration, act)
	}

	passport.ChangeValidation(newValidation)
	if act := passport.GetValidation(); act != *newValidation {
		t.Errorf("expected: %v, act: %v", newValidation, act)
	}
}

func TestPassport_String(t *testing.T) {
	const exp = "ID 77 77 777777; name Boris Goldovsky; birthday 09.04.1986; issued MVD 09.04.2010 code 770-77; registration Russia, Moscow; UFMS valid; MVD not valid; FSSP valid; document not valid;"
	passport, _ := NewPassport(
		id,
		name,
		&birthday,
		issued,
		registration,
		validation)

	if act := passport.String(); act != exp {
		t.Errorf("expected: %v, act: %v", exp, act)
	}
}
