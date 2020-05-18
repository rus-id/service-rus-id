package valuetypes

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidFirstName = errors.New("invalid first name")
	ErrInvalidLastName  = errors.New("invalid last name")
)

type Name struct {
	first  string
	middle *string
	last   string
}

func NewName(first string, middle *string, last string) (*Name, error) {
	if first == "" {
		return nil, ErrInvalidFirstName
	}

	if last == "" {
		return nil, ErrInvalidLastName
	}

	return &Name{
		first:  first,
		middle: middle,
		last:   last,
	}, nil
}

func (n *Name) GetFirst() string {
	return n.first
}

func (n *Name) GetMiddle() string {
	if n.middle == nil {
		return ""
	}

	return *n.middle
}

func (n *Name) GetLast() string {
	return n.last
}

func (n *Name) String() string {
	if n.middle == nil || *n.middle == "" {
		return fmt.Sprintf("%s %s", n.first, n.last)
	}

	return fmt.Sprintf("%s %v %s", n.first, *n.middle, n.last)
}
