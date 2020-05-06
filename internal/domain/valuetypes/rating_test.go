package valuetypes

import "testing"

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

		if rating.positive != val.positive {
			t.Errorf("expected: %v, actual: %v", val.positive, rating.positive)
		}

		if rating.negative != val.negative {
			t.Errorf("expected: %v, actual: %v", val.negative, rating.negative)
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
		rating, _ := NewRating(0, 0)

		for i := 0; i < val.positive; i++ {
			rating.AddPositive()
		}

		for i := 0; i < val.negative; i++ {
			rating.AddNegative()
		}

		if rating.positive != val.positive {
			t.Errorf("expected: %v, act: %v", val.positive, rating.positive)
		}

		if rating.negative != val.negative {
			t.Errorf("expected: %v, act: %v", val.negative, rating.negative)
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

		if rating.GetPositive() != val.positive {
			t.Errorf("expected: %v, actual: %v", val.positive, rating.positive)
		}

		if rating.GetNegative() != val.negative {
			t.Errorf("expected: %v, actual: %v", val.negative, rating.negative)
		}

		total := val.positive - val.negative
		if act := rating.GetTotal(); act != total {
			t.Errorf("expected: %v, actual: %v", total, act)
		}
	}
}
