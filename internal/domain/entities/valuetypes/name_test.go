package valuetypes

import "testing"

func TestNewName_Success(t *testing.T) {
	middle := "B"
	emptyMiddle := ""

	data := []struct {
		first  string
		middle *string
		last   string
	}{
		{"Boris", &middle, "Goldovsky"},
		{"Boris", &emptyMiddle, "Goldovsky"},
		{"Boris", nil, "Goldovsky"},
	}

	for _, val := range data {
		name, err := NewName(val.first, val.middle, val.last)

		if err != nil {
			t.Errorf("error occured: %v", err)
		}

		if name.first != val.first {
			t.Errorf("expected: %v, act: %v", val.first, name.first)
		}

		if name.middle != val.middle {
			t.Errorf("expected: %v, act: %v", val.middle, name.middle)
		}

		if name.last != val.last {
			t.Errorf("expected: %v, act: %v", val.last, name.last)
		}
	}
}

func TestNewName_Error(t *testing.T) {
	data := []struct {
		first    string
		last     string
		expected error
	}{
		{"", "Goldovsky", ErrInvalidFirstName},
		{"Boris", "", ErrInvalidLastName},
		{"", "", ErrInvalidFirstName},
	}

	for _, val := range data {
		_, err := NewName(val.first, nil, val.last)

		if err != val.expected {
			t.Errorf("expected error: %v, act: %v", val.expected, err)
		}
	}
}

func TestName_Getters(t *testing.T) {
	middle := "B"
	emptyMiddle := ""

	data := []struct {
		first          string
		middle         *string
		expectedMiddle string
		last           string
	}{
		{"Boris", &middle, middle, "Goldovsky"},
		{"Boris", &emptyMiddle, emptyMiddle, "Goldovsky"},
		{"Boris", nil, "", "Goldovsky"},
	}

	for _, val := range data {
		name, _ := NewName(val.first, val.middle, val.last)

		if name.GetFirst() != val.first {
			t.Errorf("expected: %v, act: %v", val.last, name.last)
		}

		if name.GetMiddle() != val.expectedMiddle {
			t.Errorf("expected: %v, act: %v", val.middle, name.middle)
		}

		if name.GetLast() != val.last {
			t.Errorf("expected: %v, act: %v", val.last, name.last)
		}
	}
}

func TestName_String(t *testing.T) {
	middle := "B"
	emptyMiddle := ""

	data := []struct {
		first    string
		middle   *string
		last     string
		expected string
	}{
		{"Boris", &middle, "Goldovsky", "Boris B Goldovsky"},
		{"Boris", &emptyMiddle, "Goldovsky", "Boris Goldovsky"},
		{"Boris", nil, "Goldovsky", "Boris Goldovsky"},
	}

	for _, val := range data {
		name, _ := NewName(val.first, val.middle, val.last)

		act := name.String()
		if act != val.expected {
			t.Errorf("expected: %v, act: %v", val.expected, act)
		}
	}
}
