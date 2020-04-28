package valuetypes

import (
	"errors"
	"regexp"
	"strconv"
)

var ErrInvalidSnils = errors.New("invalid snils")

var snilsRegexp = regexp.MustCompile("^[0-9]{11}$")

type Snils string

func NewSnils(value string) (Snils, error) {
	if ok, err := validateSnils(value); !ok {
		return "", err
	}

	snils := Snils(value)

	return snils, nil
}

func validateSnils(value string) (bool, error) {
	if !snilsRegexp.MatchString(value) {
		return false, ErrInvalidSnils
	}

	snils := []rune(value)
	checksum, err := checksumSnils(snils)
	if err != nil {
		return false, err
	}

	checkin := checkinSnils(checksum)

	num, err := strconv.Atoi(string(snils[9:]))
	if err != nil {
		return false, err
	}

	if checkin != num {
		return false, ErrInvalidSnils
	}

	return true, nil
}

func checkinSnils(sum int) int {
	check := 0
	if sum < 100 {
		check = sum
	} else if sum > 101 {
		check = sum % 101
		if check == 100 {
			check = 0
		}
	}

	return check
}

func checksumSnils(snils []rune) (int, error) {
	sum := 0

	for i := 0; i < 9; i++ {
		num, err := strconv.Atoi(string(snils[i]))
		if err != nil {
			return -1, err
		}
		sum += num * (9 - i)
	}

	return sum, nil
}
