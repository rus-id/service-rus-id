package user_test

import (
	"reflect"
	"testing"
	"time"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
)

func TestNewUser_Success(t *testing.T) {
	user, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		userIsRemoved,
		userVersion)

	act, _ := GetSnapshot(user, userSnapshotDate)

	if id := user.GetID(); id != *userID {
		t.Errorf("expected: %v, act: %v", *userID, id)
	}

	if act.UserID != userIDRaw {
		t.Errorf("expected: %v, act: %v", userIDRaw, act.UserID)
	}

	if act.FirstName != userFirstName {
		t.Errorf("expected: %v, act: %v", userFirstName, act.FirstName)
	}

	if act.MiddleName != userMiddleName {
		t.Errorf("expected: %v, act: %v", userMiddleName, act.MiddleName)
	}

	if act.LastName != userLastName {
		t.Errorf("expected: %v, act: %v", userLastName, act.LastName)
	}

	if act.CountryCode != userCountryCode {
		t.Errorf("expected: %v, act: %v", userCountryCode, act.UserID)
	}

	if act.PhoneNumber != userPhoneNumber {
		t.Errorf("expected: %v, act: %v", userPhoneNumber, act.PhoneNumber)
	}

	if act.RegistrationDate != userRegistrationDateStamp {
		t.Errorf("expected: %v, act: %v", userRegistrationDateStamp, act.RegistrationDate)
	}

	if act.RatingPositive != userPositiveRating {
		t.Errorf("expected: %v, act: %v", userPositiveRating, act.RatingPositive)
	}

	if act.RatingNegative != userNegativeRating {
		t.Errorf("expected: %v, act: %v", userNegativeRating, act.RatingNegative)
	}

	if !reflect.DeepEqual(act.Tolerances, userToleranceEmptySnapshots) {
		t.Errorf("expected: %v, act: %v", userToleranceEmptySnapshots, act.Tolerances)
	}

	if act.State != userStateRaw {
		t.Errorf("expected: %v, act: %v", userStateRaw, act.State)
	}

	if act.IsRemoved != userIsRemoved {
		t.Errorf("expected: %v, act: %v", userIsRemoved, act.IsRemoved)
	}

	if act.Version != userVersion {
		t.Errorf("expected: %v, act: %v", userVersion, act.Version)
	}

	if act.Timestamp != userSnapshotDateStamp {
		t.Errorf("expected: %v, act: %v", userSnapshotDateStamp, act.Timestamp)
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
			userName,
			userPhone,
			&userRegistrationDate,
			userRating,
			userState,
			ErrInvalidID,
		},
		{
			userID,
			nil,
			userPhone,
			&userRegistrationDate,
			userRating,
			userState,
			ErrInvalidName,
		},
		{
			userID,
			userName,
			nil,
			&userRegistrationDate,
			userRating,
			userState,
			ErrInvalidPhone,
		},
		{
			userID,
			userName,
			userPhone,
			nil,
			userRating,
			userState,
			ErrInvalidRegistrationDate,
		},
		{
			userID,
			userName,
			userPhone,
			&userRegistrationDate,
			nil,
			userState,
			ErrInvalidRating,
		},
		{
			userID,
			userName,
			userPhone,
			&userRegistrationDate,
			userRating,
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
			userIsRemoved,
			userVersion)

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
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		userIsRemoved,
		userVersion)

	if act := user.GetID(); act != *userID {
		t.Errorf("expected: %v, act: %v", *userID, act)
	}

	if act := user.IsRemoved(); act != userIsRemoved {
		t.Errorf("expected: %v, act: %v", userIsRemoved, act)
	}
}

func TestUser_ChangeName(t *testing.T) {
	user, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		false,
		userVersion)

	newName, _ := valuetypes.NewName("Edward", nil, "Kondratev")

	act, _ := GetSnapshot(user, userSnapshotDate)

	if exp := userName.GetFirst(); act.FirstName != exp {
		t.Errorf("expected: %v, act: %v", exp, act.FirstName)
	}

	if exp := userName.GetMiddle(); act.MiddleName != exp {
		t.Errorf("expected: %v, act: %v", exp, act.MiddleName)
	}

	if exp := userName.GetLast(); act.LastName != exp {
		t.Errorf("expected: %v, act: %v", exp, act.LastName)
	}

	user.ChangeName(newName)
	act, _ = GetSnapshot(user, userSnapshotDate)

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
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		false,
		userVersion)

	newPhone, _ := valuetypes.NewPhone(7, "9267771620")

	act, _ := GetSnapshot(user, userSnapshotDate)
	if exp := userPhone.GetCode(); valuetypes.CountryCode(act.CountryCode) != exp {
		t.Errorf("expected: %v, act: %v", exp, act.CountryCode)
	}
	if exp := userPhone.GetNumber(); act.PhoneNumber != exp {
		t.Errorf("expected: %v, act: %v", exp, act.PhoneNumber)
	}

	user.ChangePhone(newPhone)

	act, _ = GetSnapshot(user, userSnapshotDate)
	if exp := newPhone.GetCode(); valuetypes.CountryCode(act.CountryCode) != exp {
		t.Errorf("expected: %v, act: %v", exp, act.CountryCode)
	}
	if exp := newPhone.GetNumber(); act.PhoneNumber != exp {
		t.Errorf("expected: %v, act: %v", exp, act.PhoneNumber)
	}
}

func TestUser_ChangePassport(t *testing.T) {
	user, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		false,
		userVersion)

	act, _ := GetSnapshot(user, userSnapshotDate)
	if act.Passport != nil {
		t.Errorf("expected: %v, act: %v", nil, act.Passport)
	}

	user.ChangePassport(passportEntity)
	act, _ = GetSnapshot(user, userSnapshotDate)
	if !reflect.DeepEqual(act.Passport, &passportSnapshot) {
		t.Errorf("expected: %v, act: %v", passportSnapshot, act.Passport)
	}
}

func TestUser_ChangeDrivingLicense(t *testing.T) {
	user, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		false,
		userVersion)

	act, _ := GetSnapshot(user, userSnapshotDate)
	if act.DrivingLicense != nil {
		t.Errorf("expected: %v, act: %v", nil, act.DrivingLicense)
	}

	user.ChangeDrivingLicense(drivingLicenseEntity)
	act, _ = GetSnapshot(user, userSnapshotDate)
	if !reflect.DeepEqual(act.DrivingLicense, &drivingLicenseSnapshot) {
		t.Errorf("expected: %v, act: %v", drivingLicenseSnapshot, act.DrivingLicense)
	}
}

func TestUser_ChangeSnils(t *testing.T) {
	user, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		false,
		userVersion)

	act, _ := GetSnapshot(user, userSnapshotDate)
	if act.Snils != nil {
		t.Errorf("expected: %v, act: %v", nil, act.Snils)
	}

	user.ChangeSnils(userSnils)
	act, _ = GetSnapshot(user, userSnapshotDate)
	if !reflect.DeepEqual(*act.Snils, userSnilsRaw) {
		t.Errorf("expected: %v, act: %v", userSnilsRaw, *act.Snils)
	}
}

func TestUser_ChangeInn(t *testing.T) {
	user, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		false,
		userVersion)

	act, _ := GetSnapshot(user, userSnapshotDate)
	if act.Inn != nil {
		t.Errorf("expected: %v, act: %v", nil, act.Inn)
	}

	user.ChangeInn(userInn)
	act, _ = GetSnapshot(user, userSnapshotDate)
	if !reflect.DeepEqual(*act.Inn, userInnRaw) {
		t.Errorf("expected: %v, act: %v", userInnRaw, *act.Inn)
	}
}

func TestUser_ChangePhoto(t *testing.T) {
	user, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		false,
		userVersion)

	act, _ := GetSnapshot(user, userSnapshotDate)
	if act.Photo != nil {
		t.Errorf("expected: %v, act: %v", nil, act.Photo)
	}

	user.ChangePhoto(&userPhoto)
	act, _ = GetSnapshot(user, userSnapshotDate)
	if !reflect.DeepEqual(act.Photo, userPhotoRaw) {
		t.Errorf("expected: %v, act: %v", userPhotoRaw, act.Photo)
	}
}

func TestUser_ChangeCard(t *testing.T) {
	user, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		false,
		userVersion)

	act, _ := GetSnapshot(user, userSnapshotDate)
	if act.CardExpires != nil {
		t.Errorf("expected: %v, act: %v", nil, act.CardExpires)
	}
	if act.CardNumber != nil {
		t.Errorf("expected: %v, act: %v", nil, act.CardNumber)
	}

	user.ChangeCard(userCard)
	act, _ = GetSnapshot(user, userSnapshotDate)
	if *act.CardExpires != userCardExpiresStamp {
		t.Errorf("expected: %v, act: %v", userCardExpiresStamp, act.CardExpires)
	}
	if *act.CardNumber != userCardNumber {
		t.Errorf("expected: %v, act: %v", userCardNumber, act.CardNumber)
	}
}

func TestUser_Rating(t *testing.T) {
	const (
		expPositive = 10
		expNegative = 4
	)

	newRating, _ := valuetypes.NewRating(0, 0)
	user, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		newRating,
		userState,
		userIsRemoved,
		userVersion)

	act, _ := GetSnapshot(user, userSnapshotDate)
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

	act, _ = GetSnapshot(user, userSnapshotDate)
	if act.RatingPositive != expPositive {
		t.Errorf("expected: %v, act: %v", expPositive, act.RatingPositive)
	}
	if act.RatingNegative != expNegative {
		t.Errorf("expected: %v, act: %v", expNegative, act.RatingNegative)
	}
}

func TestUser_Access(t *testing.T) {
	user, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		userIsRemoved,
		userVersion)

	act, _ := GetSnapshot(user, userSnapshotDate)
	if !reflect.DeepEqual(act.Tolerances, userToleranceEmptySnapshots) {
		t.Errorf("expected: %v, act: %v", userToleranceEmptySnapshots, act.Tolerances)
	}

	expTolerance, _ := valuetypes.NewTolerance(userOtherID, []valuetypes.Accessor{valuetypes.AccessorPassport})
	expSnapshot := GetToleranceSnapshot(*userID, map[valuetypes.UserID]valuetypes.Tolerance{*userOtherID: *expTolerance})
	user.GrantAccess(*userOtherID, valuetypes.AccessorPassport)
	act, _ = GetSnapshot(user, userSnapshotDate)
	if !reflect.DeepEqual(act.Tolerances, expSnapshot) {
		t.Errorf("expected: %v, act: %v", expSnapshot, act.Tolerances)
	}

	user.RevokeAccess(*userOtherID, valuetypes.AccessorPassport)
	act, _ = GetSnapshot(user, userSnapshotDate)
	if !reflect.DeepEqual(act.Tolerances, userToleranceEmptySnapshots) {
		t.Errorf("expected: %v, act: %v", userToleranceEmptySnapshots, act.Tolerances)
	}

	user.GrantFullAccess(*userOtherID)
	act, _ = GetSnapshot(user, userSnapshotDate)
	if !reflect.DeepEqual(act.Tolerances, userToleranceSnapshots) {
		t.Errorf("expected: %v, act: %v", userToleranceSnapshots, act.Tolerances)
	}

	user.RevokeFullAccess(*userOtherID)
	act, _ = GetSnapshot(user, userSnapshotDate)
	if !reflect.DeepEqual(act.Tolerances, userToleranceEmptySnapshots) {
		t.Errorf("expected: %v, act: %v", userToleranceEmptySnapshots, act.Tolerances)
	}
}

func TestUserStates(t *testing.T) {
	user, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		userIsRemoved,
		userVersion)

	act, _ := GetSnapshot(user, userSnapshotDate)
	if valuetypes.UserState(act.State) != valuetypes.UserStateActive {
		t.Errorf("expected: %v, act: %v", valuetypes.UserStateActive, act.State)
	}

	user.Block()
	act, _ = GetSnapshot(user, userSnapshotDate)

	if valuetypes.UserState(act.State) != valuetypes.UserStateBlocked {
		t.Errorf("expected: %v, act: %v", valuetypes.UserStateBlocked, act.State)
	}

	user.Activate()
	act, _ = GetSnapshot(user, userSnapshotDate)

	if valuetypes.UserState(act.State) != valuetypes.UserStateActive {
		t.Errorf("expected: %v, act: %v", valuetypes.UserStateActive, act.State)
	}
}

func TestUser_IsRemoved(t *testing.T) {
	const expected = true

	user, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		false,
		userVersion)

	if act := user.IsRemoved(); act == expected {
		t.Errorf("expected: %v, act: %v", false, act)
	}

	user.Remove()

	if act := user.IsRemoved(); act != expected {
		t.Errorf("expected: %v, act: %v", expected, act)
	}
}

func TestUser_String(t *testing.T) {
	const expected = "User Aggregate ID 059b4e12-6983-4806-bd5a-cc3433e78f66\nName: Boris Goldovsky\nPhone: +7(903)961-53-22\nPassport: ID 77 77 777777; name Boris Goldovsky; birthday 09.04.1986; issued MVD 09.04.2010 code 770-77; registration Russia, Moscow; UFMS valid; MVD not valid; FSSP valid; document not valid;\nDriving License: ID 77 77 777777; category A; name Boris Goldovsky; birthday 09.04.1986; issued 09.04.2010; expires 09.04.2025; residence Russia; marks empty mark; GIBDD valid; document not valid;\nSNILS: 596-504-185-27\nINN: 926902267890\nPhoto: uploaded\nCard: VISA 4444333322221111 04/20\nRegistration Date: 09.03.2019 01:10:30\nRating: positive 50; negative 40; total 10;\nTolerances: ID 059b4e12-6983-4806-bd5a-cc3433e78f66 contacts, profile, phone, passport, driver license; \nState: active user\nRemoved: true\nVersion: 12345\n"
	user, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		userIsRemoved,
		userVersion)

	user.ChangeName(userName)
	user.ChangePhone(userPhone)
	user.ChangePassport(passportEntity)
	user.ChangeDrivingLicense(drivingLicenseEntity)
	user.ChangeSnils(userSnils)
	user.ChangeInn(userInn)
	user.ChangePhoto(&userPhoto)
	user.ChangeCard(userCard)
	user.Activate()
	user.GrantFullAccess(*userOtherID)

	if act := user.String(); act != expected {
		t.Errorf("expected: %v, act: %v", expected, act)
	}
}
