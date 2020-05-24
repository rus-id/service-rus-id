package valuetypes

import "errors"

type Photo []byte

var ErrInvalidPhoto = errors.New("invalid user photo")

func NewPhoto(value []byte) (Photo, error) {
	if len(value) == 0 {
		return nil, ErrInvalidPhoto
	}

	photo := Photo(value)
	return photo, nil
}

func (s Photo) String() string {
	if len(s) == 0 {
		return "Photo not exist"
	}

	return "Photo uploaded"
}
