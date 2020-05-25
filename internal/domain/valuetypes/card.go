package valuetypes

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

type PaymentSystem int

const (
	PaymentVisa PaymentSystem = iota + 1
	PaymentMaestro
	PaymentMastercard
	PaymentMir
)

var (
	ErrInvalidCardNumber = errors.New("invalid card number")

	paymentVisaRegexp       = regexp.MustCompile("^4\\d{12}(?:\\d{3})?$")
	paymentMaestroRegexp    = regexp.MustCompile("^(5018|5020|5038|6304|6759|6761|6763)[0-9]{8,15}$")
	paymentMastercardRegexp = regexp.MustCompile("^5[1-5]\\d{14}$")
	paymentMirRegexp        = regexp.MustCompile("^220\\d{13}$")
)

func GetPaymentSystem(number string) (PaymentSystem, error) {
	switch {
	case paymentVisaRegexp.MatchString(number):
		return PaymentVisa, nil
	case paymentMaestroRegexp.MatchString(number):
		return PaymentMaestro, nil
	case paymentMastercardRegexp.MatchString(number):
		return PaymentMastercard, nil
	case paymentMirRegexp.MatchString(number):
		return PaymentMir, nil
	}

	return 0, ErrInvalidCardNumber
}

func (p PaymentSystem) String() string {
	switch p {
	case PaymentVisa:
		return "VISA"
	case PaymentMaestro:
		return "Maestro"
	case PaymentMastercard:
		return "Mastercard"
	case PaymentMir:
		return "MIR"
	}

	return ""
}

type Card struct {
	number  string
	system  PaymentSystem
	expires time.Time
}

func NewCard(number string, expires time.Time) (*Card, error) {
	paymentSystem, err := GetPaymentSystem(number)
	if err != nil {
		return nil, err
	}

	return &Card{
		number:  number,
		system:  paymentSystem,
		expires: expires,
	}, nil
}

func (c Card) GetNumber() string {
	return c.number
}

func (c Card) GetLastFour() string {
	runes := []rune(c.number)
	return string(runes[len(runes)-4:])
}

func (c Card) GetExpires() time.Time {
	return c.expires
}

func (c Card) IsExpired(now time.Time) bool {
	return c.expires.Unix() < now.Unix()
}

func (c Card) String() string {
	return fmt.Sprintf("%s %s %s", c.system, c.number, c.expires.Format("01/06"))
}
