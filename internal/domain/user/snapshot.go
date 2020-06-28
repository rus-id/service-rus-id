package user

import (
	"errors"
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
	"github.com/google/uuid"
)

var (
	ErrInvalidAggregate = errors.New("user aggregate not specified")
	ErrInvalidSnapshot  = errors.New("user aggregate snapshot not specified")
)

type Snapshot struct {
	UserID           uuid.UUID `sql:",type:uuid"`
	FirstName        string
	MiddleName       *string
	LastName         string
	CountryCode      int64
	PhoneNumber      string
	Passport         *passport.Snapshot
	DrivingLicense   *driving_license.Snapshot
	Snils            *string
	Inn              *string
	Photo            []byte
	CardNumber       *string
	CardExpires      *int64
	RegistrationDate int64
	RatingPositive   int64
	RatingNegative   int64
	Tolerances       []ToleranceSnapshot
	State            int64
	IsRemoved        bool
	Version          int64
	Timestamp        int64
}

func NewSnapshot(
	userID uuid.UUID,
	firstName string,
	middleName *string,
	lastName string,
	countryCode int64,
	phoneNumber string,
	passport *passport.Snapshot,
	drivingLicense *driving_license.Snapshot,
	snils *string,
	inn *string,
	photo []byte,
	cardNumber *string,
	cardExpires *int64,
	registrationDate int64,
	ratingPositive int64,
	ratingNegative int64,
	tolerances []ToleranceSnapshot,
	state int64,
	isRemoved bool,
	version int64,
	timestamp int64,
) Snapshot {
	return Snapshot{
		UserID:           userID,
		FirstName:        firstName,
		MiddleName:       middleName,
		LastName:         lastName,
		CountryCode:      countryCode,
		PhoneNumber:      phoneNumber,
		Passport:         passport,
		DrivingLicense:   drivingLicense,
		Snils:            snils,
		Inn:              inn,
		Photo:            photo,
		CardNumber:       cardNumber,
		CardExpires:      cardExpires,
		RegistrationDate: registrationDate,
		RatingPositive:   ratingPositive,
		RatingNegative:   ratingNegative,
		Tolerances:       tolerances,
		State:            state,
		IsRemoved:        isRemoved,
		Version:          version,
		Timestamp:        timestamp,
	}
}

func GetSnapshot(user *User, timestamp time.Time) (*Snapshot, error) {
	if user == nil {
		return nil, ErrInvalidAggregate
	}

	pass := passport.GetSnapshot(user.passport)
	drivingLicense := driving_license.GetSnapshot(user.drivingLicense)
	tolerances := GetToleranceSnapshot(user.id, user.tolerances)

	var snils *string
	if user.snils != nil {
		tmp := string(*user.snils)
		snils = &tmp
	}

	var inn *string
	if user.inn != nil {
		tmp := string(*user.inn)
		inn = &tmp
	}

	var cardNumber *string
	var cardExpires *int64
	if user.card != nil {
		tmpNum := user.card.GetNumber()
		tmpExp := user.card.GetExpires().Unix()
		cardNumber = &tmpNum
		cardExpires = &tmpExp
	}

	snapshot := NewSnapshot(
		uuid.UUID(user.id),
		user.name.GetFirst(),
		user.name.GetMiddle(),
		user.name.GetLast(),
		int64(user.phone.GetCode()),
		user.phone.GetNumber(),
		pass,
		drivingLicense,
		snils,
		inn,
		user.photo,
		cardNumber,
		cardExpires,
		user.registrationDate.Unix(),
		int64(user.rating.GetPositive()),
		int64(user.rating.GetNegative()),
		tolerances,
		int64(user.state),
		user.isRemoved,
		user.version,
		timestamp.Unix())

	return &snapshot, nil
}

func LoadFromSnapshot(snapshot *Snapshot) (*User, error) {
	if snapshot == nil {
		return nil, ErrInvalidSnapshot
	}

	id, err := valuetypes.NewUserID(snapshot.UserID)
	if err != nil {
		return nil, err
	}

	name, err := valuetypes.NewName(snapshot.FirstName, snapshot.MiddleName, snapshot.LastName)
	if err != nil {
		return nil, err
	}

	phone, err := valuetypes.NewPhone(valuetypes.CountryCode(snapshot.CountryCode), snapshot.PhoneNumber)
	if err != nil {
		return nil, err
	}

	rating, err := valuetypes.NewRating(int(snapshot.RatingPositive), int(snapshot.RatingNegative))
	if err != nil {
		return nil, err
	}

	registrationDate := time.Unix(snapshot.RegistrationDate, 0).UTC()
	state := valuetypes.UserState(snapshot.State)

	user, err := NewUser(id, name, phone, &registrationDate, rating, state, snapshot.IsRemoved, snapshot.Version)
	if err != nil {
		return nil, err
	}

	user.passport, err = passport.LoadFromSnapshot(snapshot.Passport)
	if err != nil {
		return nil, err
	}

	user.drivingLicense, err = driving_license.LoadFromSnapshot(snapshot.DrivingLicense)
	if err != nil {
		return nil, err
	}

	if snapshot.Snils != nil {
		user.snils, err = valuetypes.NewSnils(*snapshot.Snils)
		if err != nil {
			return nil, err
		}
	}

	if snapshot.Inn != nil {
		user.inn, err = valuetypes.NewInn(*snapshot.Inn)
		if err != nil {
			return nil, err
		}
	}

	if snapshot.Photo != nil {
		user.photo, err = valuetypes.NewPhoto(snapshot.Photo)
		if err != nil {
			return nil, err
		}
	}

	if snapshot.CardNumber != nil && snapshot.CardExpires != nil {
		expires := time.Unix(*snapshot.CardExpires, 0).UTC()
		user.card, err = valuetypes.NewCard(*snapshot.CardNumber, expires)
		if err != nil {
			return nil, err
		}
	}

	user.tolerances, err = LoadTolerancesFromSnapshot(snapshot.Tolerances)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type ToleranceSnapshot struct {
	FromID    uuid.UUID
	ToID      uuid.UUID
	Accessors []int64
}

func NewToleranceSnapshot(
	fromID uuid.UUID,
	toID uuid.UUID,
	accessors []int64,
) ToleranceSnapshot {
	return ToleranceSnapshot{
		FromID:    fromID,
		ToID:      toID,
		Accessors: accessors,
	}
}

func GetToleranceSnapshot(userID valuetypes.UserID, m map[valuetypes.UserID]valuetypes.Tolerance) []ToleranceSnapshot {
	if m == nil {
		return nil
	}

	id := uuid.UUID(userID)
	snapshots := make([]ToleranceSnapshot, 0, len(m))

	for key, val := range m {
		toID := uuid.UUID(key)

		accessors := val.GetAccessors()
		accessorIds := make([]int64, len(accessors))
		for idx, a := range accessors {
			accessorIds[idx] = int64(a)
		}

		snapshot := NewToleranceSnapshot(id, toID, accessorIds)
		snapshots = append(snapshots, snapshot)
	}
	return snapshots
}

func LoadTolerancesFromSnapshot(snapshots []ToleranceSnapshot) (map[valuetypes.UserID]valuetypes.Tolerance, error) {
	if snapshots == nil {
		return make(map[valuetypes.UserID]valuetypes.Tolerance), nil
	}

	tolerances := make(map[valuetypes.UserID]valuetypes.Tolerance, len(snapshots))

	for _, snapshot := range snapshots {

		accessors := make([]valuetypes.Accessor, len(snapshot.Accessors))
		for idx, a := range snapshot.Accessors {
			accessors[idx] = valuetypes.Accessor(a)
		}

		userID := valuetypes.UserID(snapshot.ToID)
		tolerance, err := valuetypes.NewTolerance(&userID, accessors)
		if err != nil {
			return nil, err
		}

		tolerances[userID] = *tolerance
	}

	return tolerances, nil
}
