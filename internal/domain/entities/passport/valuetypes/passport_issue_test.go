package valuetypes_test

import (
	"testing"
	"time"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport/valuetypes"
)

func TestNewPassportIssue(t *testing.T) {
	date := time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	data := []struct {
		organisation string
		date         time.Time
		code         string
		err          error
	}{
		{"УФМС ХХХ", date, "770-04", nil},
		{"", date, "770-04", ErrInvalidOrganisation},
		{"УФМС ХХХ", date, "", ErrInvalidCode},
		{"", date, "", ErrInvalidOrganisation},
	}

	for _, val := range data {
		issue, err := NewPassportIssue(val.organisation, val.date, val.code)

		if err != val.err {
			t.Errorf("expected err: %v, issue: %v", val.err, err)
		}

		if err != nil {
			continue
		}

		if act := issue.GetOrganisation(); act != val.organisation {
			t.Errorf("expected: %v, actual: %v", val.organisation, act)
		}

		if act := issue.GetDate(); act != val.date {
			t.Errorf("expected: %v, actual: %v", val.date, act)
		}

		if act := issue.GetCode(); act != val.code {
			t.Errorf("expected: %v, actual: %v", val.code, act)
		}
	}
}

func TestPassportIssue_Getters(t *testing.T) {
	date := time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	data := []struct {
		organisation string
		date         time.Time
		code         string
	}{
		{"УФМС ХХХ", date, "770-04"},
		{"УФМС YYY", date, "QWERTY"},
	}

	for _, val := range data {
		issue, _ := NewPassportIssue(val.organisation, val.date, val.code)

		if act := issue.GetOrganisation(); act != val.organisation {
			t.Errorf("expected: %v, actual: %v", val.organisation, act)
		}

		if act := issue.GetDate(); act != val.date {
			t.Errorf("expected: %v, actual: %v", val.date, act)
		}

		if act := issue.GetCode(); act != val.code {
			t.Errorf("expected: %v, actual: %v", val.code, act)
		}
	}
}

func TestPassportIssue_String(t *testing.T) {
	date := time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	const expected = "Issued by УФМС ХХХ, 09.04.2010, code 770-04"

	issue, _ := NewPassportIssue("УФМС ХХХ", date, "770-04")

	if act := issue.String(); act != expected {
		t.Errorf("expected: %v, actual: %v", expected, act)
	}
}
