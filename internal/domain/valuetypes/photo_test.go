package valuetypes_test

import (
	"reflect"
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestNewPhoto_Success(t *testing.T) {
	raw := []byte{10, 20, 30, 40, 50}
	exp := Photo(raw)

	photo, err := NewPhoto(raw)

	if err != nil {
		t.Errorf("error occured: %v", err)
	}

	if !reflect.DeepEqual(photo, exp) {
		t.Errorf("expected: %v, actual: %v", exp, photo)
	}
}

func TestNewPhoto_Error(t *testing.T) {
	raw := make([]byte, 0)
	_, err := NewPhoto(raw)

	if err != ErrInvalidPhoto {
		t.Errorf("expected error: %v, actual: %v", ErrInvalidPhoto, err)
	}
}

func TestPhoto_String(t *testing.T) {
	data := []struct {
		photo Photo
		exp   string
	}{
		{[]byte{10, 20, 30, 40, 50}, "uploaded"},
		{[]byte{}, "not uploaded"},
	}

	for _, val := range data {
		if act := val.photo.String(); act != val.exp {
			t.Errorf("expected: %q, act: %v", val.exp, act)
		}
	}
}
