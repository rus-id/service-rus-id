package valuetypes

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidPositiveRate = errors.New("invalid negative rate")
	ErrInvalidNegativeRate = errors.New("invalid negative rate")
)

type Rating struct {
	negative int
	positive int
}

func NewRating(positive, negative int) (*Rating, error) {
	if positive < 0 {
		return nil, ErrInvalidPositiveRate
	}

	if negative < 0 {
		return nil, ErrInvalidNegativeRate
	}

	return &Rating{negative: negative, positive: positive}, nil
}

func (r Rating) AddPositive() Rating {
	rating, _ := NewRating(r.positive+1, r.negative)
	return *rating
}

func (r Rating) AddNegative() Rating {
	rating, _ := NewRating(r.positive, r.negative+1)
	return *rating
}

func (r Rating) GetPositive() int {
	return r.positive
}

func (r Rating) GetNegative() int {
	return r.negative
}

func (r Rating) GetTotal() int {
	return r.positive - r.negative
}

func (r Rating) String() string {
	return fmt.Sprintf("positive %d; negative %d; total %d;", r.positive, r.negative, r.GetTotal())
}
