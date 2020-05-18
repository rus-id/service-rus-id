package valuetypes

import (
	"errors"
	"regexp"
	"strconv"
)

var ErrInvalidInn = errors.New("invalid inn")

var innRegexp = regexp.MustCompile("^[0-9]+$")

type Inn string

func NewInn(value string) (Inn, error) {
	if ok, err := ValidateInn(value); !ok {
		return "", err
	}

	return Inn(value), nil
}

func ValidateInn(value string) (bool, error) {
	if !innRegexp.MatchString(value) {
		return false, ErrInvalidInn
	}

	inn := []rune(value)
	length := len(inn)

	if length == 10 {
		coefficients10 := []int{2, 4, 10, 3, 5, 9, 4, 6, 8}
		return checkInn(inn, coefficients10, length)
	}

	if length == 12 {
		coefficients11 := []int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
		coefficients12 := []int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

		isValid11, err := checkInn(inn, coefficients11, length-1)
		if err != nil {
			return false, err
		}

		isValid12, err := checkInn(inn, coefficients12, length)
		if err != nil {
			return false, err
		}

		return isValid11 && isValid12, nil
	}

	return false, ErrInvalidInn
}

func checkInn(inn []rune, coefficients []int, length int) (bool, error) {
	checksum, err := checksumInn(inn, coefficients)
	if err != nil {
		return false, err
	}

	lastNum, err := strconv.Atoi(string(inn[length-1]))
	if err != nil {
		return false, err
	}

	if checksum != lastNum {
		return false, ErrInvalidInn
	}

	return true, nil
}

func checksumInn(inn []rune, coefficients []int) (int, error) {
	num := 0
	for idx, val := range coefficients {
		innNum, err := strconv.Atoi(string(inn[idx]))
		if err != nil {
			return 0, err
		}
		num += val * innNum
	}

	return num % 11 % 10, nil
}
