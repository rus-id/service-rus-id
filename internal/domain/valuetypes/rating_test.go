package valuetypes_test

import (
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestNewRating(t *testing.T) {
	data := []struct {
		positive int
		negative int
		err      error
	}{
		{10, 10, nil},
		{0, 10, nil},
		{10, 0, nil},
		{0, 0, nil},
		{-10, 10, ErrInvalidPositiveRate},
		{10, -10, ErrInvalidNegativeRate},
	}

	for _, val := range data {
		rating, err := NewRating(val.positive, val.negative)
		if err != val.err {
			t.Errorf("expected error: %v, actual: %v", val.err, err)
		}

		if err != nil {
			continue
		}

		if act := rating.GetPositive(); act != val.positive {
			t.Errorf("expected: %v, actual: %v", val.positive, act)
		}

		if act := rating.GetNegative(); act != val.negative {
			t.Errorf("expected: %v, actual: %v", val.negative, act)
		}
	}
}

func TestRating_Setters(t *testing.T) {
	data := []struct {
		positive int
		negative int
		total    int
	}{
		{1, 1, 0},
		{5, 4, 1},
		{5, 6, -1},
		{0, 0, 0},
	}

	for _, val := range data {
		newRating, _ := NewRating(0, 0)

		rating := *newRating
		for i := 0; i < val.positive; i++ {
			rating = rating.AddPositive()
		}

		for i := 0; i < val.negative; i++ {
			rating = rating.AddNegative()
		}

		if act := rating.GetPositive(); act != val.positive {
			t.Errorf("expected: %v, act: %v", val.positive, act)
		}

		if act := rating.GetNegative(); act != val.negative {
			t.Errorf("expected: %v, act: %v", val.negative, act)
		}

		if total := rating.GetTotal(); total != val.total {
			t.Errorf("expected: %v, act: %v", val.total, total)
		}
	}
}

func TestRating_Getters(t *testing.T) {
	data := []struct {
		positive int
		negative int
	}{
		{10, 10},
		{0, 10},
		{10, 0},
		{0, 0},
	}

	for _, val := range data {
		rating, _ := NewRating(val.positive, val.negative)

		if act := rating.GetPositive(); act != val.positive {
			t.Errorf("expected: %v, actual: %v", val.positive, act)
		}

		if act := rating.GetNegative(); act != val.negative {
			t.Errorf("expected: %v, actual: %v", val.negative, act)
		}

		total := val.positive - val.negative
		if act := rating.GetTotal(); act != total {
			t.Errorf("expected: %v, actual: %v", total, act)
		}
	}
}
