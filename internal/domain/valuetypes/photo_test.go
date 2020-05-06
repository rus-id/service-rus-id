package valuetypes

import (
	"reflect"
	"testing"
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
