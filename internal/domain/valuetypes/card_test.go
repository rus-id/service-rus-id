package valuetypes

import (
	"testing"
	"time"
)

var paymentSystemData = []struct {
	number string
	system PaymentSystem
	text   string
}{
	{"4444333322221111", PaymentVisa, "VISA"},
	{"6759649826438453", PaymentMaestro, "Maestro"},
	{"5555555555554444", PaymentMastercard, "Mastercard"},
	{"2201382000000013", PaymentMir, "MIR"},
	{"7701382000000013", 0, ""},
}

func TestGetPaymentSystem_Success(t *testing.T) {
	for _, val := range paymentSystemData {
		system, _ := getPaymentSystem(val.number)
		if val.system != system {
			t.Errorf("expected: %v, actual: %v", val.system, system)
		}
	}
}

func TestGetPaymentSystem_Fail(t *testing.T) {
	_, err := getPaymentSystem("1234567890")
	if err != ErrInvalidCardNumber {
		t.Errorf("expected error: %v, actual: %v", ErrInvalidCardNumber, err)
	}
}

func TestPaymentSystem_String(t *testing.T) {
	for _, val := range paymentSystemData {
		system, _ := getPaymentSystem(val.number)
		if text := system.String(); text != val.text {
			t.Errorf("expected: %v, actual: %v", val.text, text)
		}
	}
}

func TestNewCard(t *testing.T) {
	expired := time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC)

	data := []struct {
		number string
		err    error
	}{
		{"4444333322221111", nil},
		{"6759649826438453", nil},
		{"5555555555554444", nil},
		{"2201382000000013", nil},
		{"555550013", ErrInvalidCardNumber},
		{"", ErrInvalidCardNumber},
	}

	for _, val := range data {
		card, err := NewCard(val.number, expired)

		if err != val.err {
			t.Errorf("expected error: %v, actual: %v", val.err, err)
		}

		if err != nil {
			continue
		}

		if card.number != val.number {
			t.Errorf("expected: %v, actual: %v", val.number, card.number)
		}

		if card.expired != expired {
			t.Errorf("expected: %v, actual: %v", expired, card.expired)
		}
	}
}

func TestCard_IsExpired(t *testing.T) {
	data := []struct {
		expired   time.Time
		now       time.Time
		isExpired bool
	}{
		{
			time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC),
			time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC),
			false,
		},
		{
			time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC),
			time.Date(2020, time.Month(5), 9, 1, 10, 30, 0, time.UTC),
			true,
		},
		{
			time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC),
			time.Date(2020, time.Month(3), 9, 1, 10, 30, 0, time.UTC),
			false,
		},
	}

	number := "4444333322221111"
	for _, val := range data {
		card, _ := NewCard(number, val.expired)
		if isExpired := card.IsExpired(val.now); isExpired != val.isExpired {
			t.Errorf("expected: %v, actual: %v", val.isExpired, isExpired)
		}
	}
}

func TestCard_Getters(t *testing.T) {
	expired := time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	data := []struct {
		number   string
		lastFour string
	}{
		{"4444333322221111", "1111"},
		{"6759649826438453", "8453"},
		{"5555555555554444", "4444"},
		{"2201382000000013", "0013"},
	}

	for _, val := range data {
		card, _ := NewCard(val.number, expired)
		expired := time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
		if act := card.GetNumber(); act != val.number {
			t.Errorf("expected: %v, actual: %v", val.number, act)
		}

		if act := card.GetLastFour(); act != val.lastFour {
			t.Errorf("expected: %v, actual: %v", val.lastFour, act)
		}

		if act := card.GetExpired(); act != expired {
			t.Errorf("expected: %v, actual: %v", expired, act)
		}
	}
}

func TestCard_String(t *testing.T) {
	expired := time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	card, _ := NewCard("4444333322221111", expired)

	expected := "VISA 4444333322221111. Expired 04/20."

	if act := card.String(); act != expected {
		t.Errorf("expected: %v, actual: %v", expired, act)
	}
}