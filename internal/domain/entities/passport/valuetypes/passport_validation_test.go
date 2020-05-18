package valuetypes

import "testing"

func TestPassportValidation_Getters(t *testing.T) {
	data := []struct {
		ufms     bool
		mvd      bool
		fssp     bool
		document bool
	}{
		{false, false, false, false},
		{true, false, false, false},
		{false, true, false, false},
		{false, false, true, false},
		{false, false, false, true},
		{true, false, true, true},
		{true, true, true, true},
	}

	for _, val := range data {
		validation := NewPassportValidation(val.ufms, val.mvd, val.fssp, val.document)

		if act := validation.GetUfms(); act != val.ufms {
			t.Errorf("expected: %v, act: %v", val.ufms, act)
		}

		if act := validation.GetMvd(); act != val.mvd {
			t.Errorf("expected: %v, act: %v", val.mvd, act)
		}

		if act := validation.GetFssp(); act != val.fssp {
			t.Errorf("expected: %v, act: %v", val.fssp, act)
		}

		if act := validation.GetDocument(); act != val.document {
			t.Errorf("expected: %v, act: %v", val.document, act)
		}
	}
}

func TestPassport_String(t *testing.T) {
	const expected = "UFMS valid: true. MVD valid: false. FSSP valid: true Document valid: false."

	act := NewPassportValidation(true, false, true, false).String()

	if act != expected {
		t.Errorf("expected text: %s, act: %s", expected, act)
	}
}
