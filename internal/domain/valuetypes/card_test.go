package valuetypes_test

import (
	"testing"
	"time"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
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
		system, _ := GetPaymentSystem(val.number)
		if val.system != system {
			t.Errorf("expected: %v, actual: %v", val.system, system)
		}
	}
}

func TestGetPaymentSystem_Fail(t *testing.T) {
	_, err := GetPaymentSystem("1234567890")
	if err != ErrInvalidCardNumber {
		t.Errorf("expected error: %v, actual: %v", ErrInvalidCardNumber, err)
	}
}

func TestPaymentSystem_String(t *testing.T) {
	for _, val := range paymentSystemData {
		system, _ := GetPaymentSystem(val.number)
		if text := system.String(); text != val.text {
			t.Errorf("expected: %q, actual: %q", val.text, text)
		}
	}
}

func TestNewCard(t *testing.T) {
	expires := time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC)

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
		card, err := NewCard(val.number, expires)

		if err != val.err {
			t.Errorf("expected error: %v, actual: %v", val.err, err)
		}

		if err != nil {
			continue
		}

		if act := card.GetNumber(); act != val.number {
			t.Errorf("expected: %v, actual: %v", val.number, act)
		}

		if act := card.GetExpires(); act != expires {
			t.Errorf("expected: %v, actual: %v", expires, act)
		}
	}
}

func TestCard_IsExpired(t *testing.T) {
	data := []struct {
		expires   time.Time
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
		card, _ := NewCard(number, val.expires)
		if isExpired := card.IsExpired(val.now); isExpired != val.isExpired {
			t.Errorf("expected: %v, actual: %v", val.isExpired, isExpired)
		}
	}
}

func TestCard_Getters(t *testing.T) {
	expires := time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
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
		card, _ := NewCard(val.number, expires)
		expires := time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
		if act := card.GetNumber(); act != val.number {
			t.Errorf("expected: %v, actual: %v", val.number, act)
		}

		if act := card.GetLastFour(); act != val.lastFour {
			t.Errorf("expected: %v, actual: %v", val.lastFour, act)
		}

		if act := card.GetExpires(); act != expires {
			t.Errorf("expected: %v, actual: %v", expires, act)
		}
	}
}

func TestCard_String(t *testing.T) {
	expires := time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	card, _ := NewCard("4444333322221111", expires)

	expected := "VISA 4444333322221111 04/20"

	if act := card.String(); act != expected {
		t.Errorf("expected: %q, actual: %q", expires, act)
	}
}
