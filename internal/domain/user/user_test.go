package user_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/mock"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestNewUser_Success(t *testing.T) {
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		mock.UserIsRemoved,
		mock.UserVersion)

	act, _ := GetSnapshot(user, mock.UserSnapshotDate)

	if id := user.GetID(); id != *mock.UserID {
		t.Errorf("expected: %v, act: %v", *mock.UserID, id)
	}

	if act.UserID != mock.UserIDRaw {
		t.Errorf("expected: %v, act: %v", mock.UserIDRaw, act.UserID)
	}

	if act.FirstName != mock.UserFirstName {
		t.Errorf("expected: %v, act: %v", mock.UserFirstName, act.FirstName)
	}

	if act.MiddleName != mock.UserMiddleName {
		t.Errorf("expected: %v, act: %v", mock.UserMiddleName, act.MiddleName)
	}

	if act.LastName != mock.UserLastName {
		t.Errorf("expected: %v, act: %v", mock.UserLastName, act.LastName)
	}

	if act.CountryCode != mock.UserCountryCode {
		t.Errorf("expected: %v, act: %v", mock.UserCountryCode, act.UserID)
	}

	if act.PhoneNumber != mock.UserPhoneNumber {
		t.Errorf("expected: %v, act: %v", mock.UserPhoneNumber, act.PhoneNumber)
	}

	if act.RegistrationDate != mock.UserRegistrationDateStamp {
		t.Errorf("expected: %v, act: %v", mock.UserRegistrationDateStamp, act.RegistrationDate)
	}

	if act.RatingPositive != mock.UserPositiveRating {
		t.Errorf("expected: %v, act: %v", mock.UserPositiveRating, act.RatingPositive)
	}

	if act.RatingNegative != mock.UserNegativeRating {
		t.Errorf("expected: %v, act: %v", mock.UserNegativeRating, act.RatingNegative)
	}

	if !reflect.DeepEqual(act.Tolerances, mock.UserToleranceEmptySnapshots) {
		t.Errorf("expected: %v, act: %v", mock.UserToleranceEmptySnapshots, act.Tolerances)
	}

	if act.State != mock.UserStateRaw {
		t.Errorf("expected: %v, act: %v", mock.UserStateRaw, act.State)
	}

	if act.IsRemoved != mock.UserIsRemoved {
		t.Errorf("expected: %v, act: %v", mock.UserIsRemoved, act.IsRemoved)
	}

	if act.Version != mock.UserVersion {
		t.Errorf("expected: %v, act: %v", mock.UserVersion, act.Version)
	}

	if act.Timestamp != mock.UserSnapshotDateStamp {
		t.Errorf("expected: %v, act: %v", mock.UserSnapshotDateStamp, act.Timestamp)
	}
}

func TestNewUser_Error(t *testing.T) {
	data := []struct {
		id               *valuetypes.UserID
		name             *valuetypes.Name
		phone            *valuetypes.Phone
		registrationDate *time.Time
		rating           *valuetypes.Rating
		state            valuetypes.UserState
		err              error
	}{
		{
			nil,
			mock.UserName,
			mock.UserPhone,
			&mock.UserRegistrationDate,
			mock.UserRating,
			mock.UserState,
			ErrInvalidID,
		},
		{
			mock.UserID,
			nil,
			mock.UserPhone,
			&mock.UserRegistrationDate,
			mock.UserRating,
			mock.UserState,
			ErrInvalidName,
		},
		{
			mock.UserID,
			mock.UserName,
			nil,
			&mock.UserRegistrationDate,
			mock.UserRating,
			mock.UserState,
			ErrInvalidPhone,
		},
		{
			mock.UserID,
			mock.UserName,
			mock.UserPhone,
			nil,
			mock.UserRating,
			mock.UserState,
			ErrInvalidRegistrationDate,
		},
		{
			mock.UserID,
			mock.UserName,
			mock.UserPhone,
			&mock.UserRegistrationDate,
			nil,
			mock.UserState,
			ErrInvalidRating,
		},
		{
			mock.UserID,
			mock.UserName,
			mock.UserPhone,
			&mock.UserRegistrationDate,
			mock.UserRating,
			100,
			ErrInvalidState,
		},
	}

	for _, val := range data {
		user, err := NewUser(
			val.id,
			val.name,
			val.phone,
			val.registrationDate,
			val.rating,
			val.state,
			mock.UserIsRemoved,
			mock.UserVersion)

		if user != nil {
			t.Errorf("expected: %v, act: %v", nil, user)
		}

		if err != val.err {
			t.Errorf("expected: %v, act: %v", val.err, err)
		}
	}
}

func TestUser_Getters(t *testing.T) {
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		mock.UserIsRemoved,
		mock.UserVersion)

	if act := user.GetID(); act != *mock.UserID {
		t.Errorf("expected: %v, act: %v", *mock.UserID, act)
	}

	if act := user.IsRemoved(); act != mock.UserIsRemoved {
		t.Errorf("expected: %v, act: %v", mock.UserIsRemoved, act)
	}
}

func TestUser_ChangeName(t *testing.T) {
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		false,
		mock.UserVersion)

	newName, _ := valuetypes.NewName("Edward", nil, "Kondratev")

	act, _ := GetSnapshot(user, mock.UserSnapshotDate)

	if exp := mock.UserName.GetFirst(); act.FirstName != exp {
		t.Errorf("expected: %v, act: %v", exp, act.FirstName)
	}

	if exp := mock.UserName.GetMiddle(); act.MiddleName != exp {
		t.Errorf("expected: %v, act: %v", exp, act.MiddleName)
	}

	if exp := mock.UserName.GetLast(); act.LastName != exp {
		t.Errorf("expected: %v, act: %v", exp, act.LastName)
	}

	user.ChangeName(newName)
	act, _ = GetSnapshot(user, mock.UserSnapshotDate)

	if exp := newName.GetFirst(); act.FirstName != exp {
		t.Errorf("expected: %v, act: %v", exp, act.FirstName)
	}

	if exp := newName.GetMiddle(); act.MiddleName != exp {
		t.Errorf("expected: %v, act: %v", exp, act.MiddleName)
	}

	if exp := newName.GetLast(); act.LastName != exp {
		t.Errorf("expected: %v, act: %v", exp, act.LastName)
	}

	user.ChangeName(newName)
}

func TestUser_ChangePhone(t *testing.T) {
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		false,
		mock.UserVersion)

	newPhone, _ := valuetypes.NewPhone(7, "9267771620")

	act, _ := GetSnapshot(user, mock.UserSnapshotDate)
	if exp := mock.UserPhone.GetCode(); valuetypes.CountryCode(act.CountryCode) != exp {
		t.Errorf("expected: %v, act: %v", exp, act.CountryCode)
	}
	if exp := mock.UserPhone.GetNumber(); act.PhoneNumber != exp {
		t.Errorf("expected: %v, act: %v", exp, act.PhoneNumber)
	}

	user.ChangePhone(newPhone)

	act, _ = GetSnapshot(user, mock.UserSnapshotDate)
	if exp := newPhone.GetCode(); valuetypes.CountryCode(act.CountryCode) != exp {
		t.Errorf("expected: %v, act: %v", exp, act.CountryCode)
	}
	if exp := newPhone.GetNumber(); act.PhoneNumber != exp {
		t.Errorf("expected: %v, act: %v", exp, act.PhoneNumber)
	}
}

func TestUser_ChangePassport(t *testing.T) {
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		false,
		mock.UserVersion)

	act, _ := GetSnapshot(user, mock.UserSnapshotDate)
	if act.Passport != nil {
		t.Errorf("expected: %v, act: %v", nil, act.Passport)
	}

	user.ChangePassport(mock.PassportEntity)
	act, _ = GetSnapshot(user, mock.UserSnapshotDate)
	if !reflect.DeepEqual(act.Passport, &mock.PassportSnapshot) {
		t.Errorf("expected: %v, act: %v", mock.PassportSnapshot, act.Passport)
	}
}

func TestUser_ChangeDrivingLicense(t *testing.T) {
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		false,
		mock.UserVersion)

	act, _ := GetSnapshot(user, mock.UserSnapshotDate)
	if act.DrivingLicense != nil {
		t.Errorf("expected: %v, act: %v", nil, act.DrivingLicense)
	}

	user.ChangeDrivingLicense(mock.DrivingLicenseEntity)
	act, _ = GetSnapshot(user, mock.UserSnapshotDate)
	if !reflect.DeepEqual(act.DrivingLicense, &mock.DrivingLicenseSnapshot) {
		t.Errorf("expected: %v, act: %v", mock.DrivingLicenseSnapshot, act.DrivingLicense)
	}
}

func TestUser_ChangeSnils(t *testing.T) {
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		false,
		mock.UserVersion)

	act, _ := GetSnapshot(user, mock.UserSnapshotDate)
	if act.Snils != nil {
		t.Errorf("expected: %v, act: %v", nil, act.Snils)
	}

	user.ChangeSnils(mock.UserSnils)
	act, _ = GetSnapshot(user, mock.UserSnapshotDate)
	if !reflect.DeepEqual(*act.Snils, mock.UserSnilsRaw) {
		t.Errorf("expected: %v, act: %v", mock.UserSnilsRaw, *act.Snils)
	}
}

func TestUser_ChangeInn(t *testing.T) {
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		false,
		mock.UserVersion)

	act, _ := GetSnapshot(user, mock.UserSnapshotDate)
	if act.Inn != nil {
		t.Errorf("expected: %v, act: %v", nil, act.Inn)
	}

	user.ChangeInn(mock.UserInn)
	act, _ = GetSnapshot(user, mock.UserSnapshotDate)
	if !reflect.DeepEqual(*act.Inn, mock.UserInnRaw) {
		t.Errorf("expected: %v, act: %v", mock.UserInnRaw, *act.Inn)
	}
}

func TestUser_ChangePhoto(t *testing.T) {
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		false,
		mock.UserVersion)

	act, _ := GetSnapshot(user, mock.UserSnapshotDate)
	if act.Photo != nil {
		t.Errorf("expected: %v, act: %v", nil, act.Photo)
	}

	user.ChangePhoto(&mock.UserPhoto)
	act, _ = GetSnapshot(user, mock.UserSnapshotDate)
	if !reflect.DeepEqual(act.Photo, mock.UserPhotoRaw) {
		t.Errorf("expected: %v, act: %v", mock.UserPhotoRaw, act.Photo)
	}
}

func TestUser_ChangeCard(t *testing.T) {
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		false,
		mock.UserVersion)

	act, _ := GetSnapshot(user, mock.UserSnapshotDate)
	if act.CardExpires != nil {
		t.Errorf("expected: %v, act: %v", nil, act.CardExpires)
	}
	if act.CardNumber != nil {
		t.Errorf("expected: %v, act: %v", nil, act.CardNumber)
	}

	user.ChangeCard(mock.UserCard)
	act, _ = GetSnapshot(user, mock.UserSnapshotDate)
	if *act.CardExpires != mock.UserCardExpiresStamp {
		t.Errorf("expected: %v, act: %v", mock.UserCardExpiresStamp, act.CardExpires)
	}
	if *act.CardNumber != mock.UserCardNumber {
		t.Errorf("expected: %v, act: %v", mock.UserCardNumber, act.CardNumber)
	}
}

func TestUser_Rating(t *testing.T) {
	const (
		expPositive = 10
		expNegative = 4
	)

	newRating, _ := valuetypes.NewRating(0, 0)
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		newRating,
		mock.UserState,
		false,
		mock.UserVersion)

	act, _ := GetSnapshot(user, mock.UserSnapshotDate)
	if act.RatingPositive != 0 {
		t.Errorf("expected: %v, act: %v", 0, act.RatingPositive)
	}
	if act.RatingNegative != 0 {
		t.Errorf("expected: %v, act: %v", 0, act.RatingNegative)
	}

	for i := 0; i < expPositive; i++ {
		user.IncreaseRating()
	}

	for i := 0; i < expNegative; i++ {
		user.DecreaseRating()
	}

	act, _ = GetSnapshot(user, mock.UserSnapshotDate)
	if act.RatingPositive != expPositive {
		t.Errorf("expected: %v, act: %v", expPositive, act.RatingPositive)
	}
	if act.RatingNegative != expNegative {
		t.Errorf("expected: %v, act: %v", expNegative, act.RatingNegative)
	}
}

func TestUser_Access(t *testing.T) {
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		false,
		mock.UserVersion)

	act, _ := GetSnapshot(user, mock.UserSnapshotDate)
	if !reflect.DeepEqual(act.Tolerances, mock.UserToleranceEmptySnapshots) {
		t.Errorf("expected: %v, act: %v", mock.UserToleranceEmptySnapshots, act.Tolerances)
	}

	expTolerance, _ := valuetypes.NewTolerance(mock.UserOtherID, []valuetypes.Accessor{valuetypes.AccessorPassport})
	expSnapshot := GetToleranceSnapshot(*mock.UserID, map[valuetypes.UserID]valuetypes.Tolerance{*mock.UserOtherID: *expTolerance})
	user.GrantAccess(*mock.UserOtherID, valuetypes.AccessorPassport)
	act, _ = GetSnapshot(user, mock.UserSnapshotDate)
	if !reflect.DeepEqual(act.Tolerances, expSnapshot) {
		t.Errorf("expected: %v, act: %v", expSnapshot, act.Tolerances)
	}

	user.RevokeAccess(*mock.UserOtherID, valuetypes.AccessorPassport)
	act, _ = GetSnapshot(user, mock.UserSnapshotDate)
	if !reflect.DeepEqual(act.Tolerances, mock.UserToleranceEmptySnapshots) {
		t.Errorf("expected: %v, act: %v", mock.UserToleranceEmptySnapshots, act.Tolerances)
	}

	user.GrantFullAccess(*mock.UserOtherID)
	act, _ = GetSnapshot(user, mock.UserSnapshotDate)
	if !reflect.DeepEqual(act.Tolerances, mock.UserToleranceSnapshots) {
		t.Errorf("expected: %v, act: %v", mock.UserToleranceSnapshots, act.Tolerances)
	}

	user.RevokeFullAccess(*mock.UserOtherID)
	act, _ = GetSnapshot(user, mock.UserSnapshotDate)
	if !reflect.DeepEqual(act.Tolerances, mock.UserToleranceEmptySnapshots) {
		t.Errorf("expected: %v, act: %v", mock.UserToleranceEmptySnapshots, act.Tolerances)
	}
}

func TestUserStates(t *testing.T) {
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		false,
		mock.UserVersion)

	act, _ := GetSnapshot(user, mock.UserSnapshotDate)
	if valuetypes.UserState(act.State) != valuetypes.UserStateActive {
		t.Errorf("expected: %v, act: %v", valuetypes.UserStateActive, act.State)
	}

	user.Block()
	act, _ = GetSnapshot(user, mock.UserSnapshotDate)

	if valuetypes.UserState(act.State) != valuetypes.UserStateBlocked {
		t.Errorf("expected: %v, act: %v", valuetypes.UserStateBlocked, act.State)
	}

	user.Activate()
	act, _ = GetSnapshot(user, mock.UserSnapshotDate)

	if valuetypes.UserState(act.State) != valuetypes.UserStateActive {
		t.Errorf("expected: %v, act: %v", valuetypes.UserStateActive, act.State)
	}
}

func TestUser_IsRemoved(t *testing.T) {
	const expected = true

	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		false,
		mock.UserVersion)

	if act := user.IsRemoved(); act == expected {
		t.Errorf("expected: %v, act: %v", false, act)
	}

	user.Remove()

	if act := user.IsRemoved(); act != expected {
		t.Errorf("expected: %v, act: %v", expected, act)
	}
}

func TestUser_String(t *testing.T) {
	const expected = "User Aggregate ID 059b4e12-6983-4806-bd5a-cc3433e78f66\nName: Boris Goldovsky\nPhone: +7(903)961-53-22\nPassport: ID 77 77 777777; name Boris Goldovsky; birthday 09.04.1986; issued MVD 09.04.2010 code 770-77; registration Russia, Moscow; UFMS valid; MVD not valid; FSSP valid; document not valid;\nDriving License: ID 77 77 777777; category A; name Boris Goldovsky; birthday 09.04.1986; issued 09.04.2010; expires 09.04.2025; residence Russia; marks empty mark; GIBDD valid; document not valid;\nSNILS: 596-504-185-27\nINN: 926902267890\nPhoto: uploaded\nCard: VISA 4444333322221111 04/20\nRegistration Date: 09.03.2019 01:10:30\nRating: positive 50; negative 40; total 10;\nTolerances: ID 060b4e12-6983-4806-bd5a-cc3433e78f66 contacts, profile, phone, passport, driver license; \nState: active user\nRemoved: true\nVersion: 12345\n"
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		false,
		mock.UserVersion)

	user.ChangeName(mock.UserName)
	user.ChangePhone(mock.UserPhone)
	user.ChangePassport(mock.PassportEntity)
	user.ChangeDrivingLicense(mock.DrivingLicenseEntity)
	user.ChangeSnils(mock.UserSnils)
	user.ChangeInn(mock.UserInn)
	user.ChangePhoto(&mock.UserPhoto)
	user.ChangeCard(mock.UserCard)
	user.Activate()
	user.GrantFullAccess(*mock.UserOtherID)
	user.Remove()

	if act := user.String(); act != expected {
		t.Errorf("expected: %v, act: %v", expected, act)
	}
}
