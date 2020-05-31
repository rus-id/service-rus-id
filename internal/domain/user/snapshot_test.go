package user_test

import (
	"reflect"
	"testing"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/user"
)

func TestNewSnapshot(t *testing.T) {
	act := NewSnapshot(
		userIDRaw,
		userFirstName,
		userMiddleName,
		userLastName,
		userCountryCode,
		userPhoneNumber,
		&passportSnapshot,
		&drivingLicenseSnapshot,
		&userSnilsRaw,
		&userInnRaw,
		userPhotoRaw,
		&userCardNumber,
		&userCardExpiresStamp,
		userRegistrationDateStamp,
		userPositiveRating,
		userNegativeRating,
		userToleranceSnapshots,
		userStateRaw,
		userIsRemoved,
		userVersion,
		userSnapshotDateStamp,
	)

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

	if !reflect.DeepEqual(act.Passport, &passportSnapshot) {
		t.Errorf("expected: %v, act: %v", &passportSnapshot, act.Passport)
	}

	if !reflect.DeepEqual(act.DrivingLicense, &drivingLicenseSnapshot) {
		t.Errorf("expected: %v, act: %v", &drivingLicenseSnapshot, act.DrivingLicense)
	}

	if *act.Snils != userSnilsRaw {
		t.Errorf("expected: %v, act: %v", userSnilsRaw, *act.Snils)
	}

	if *act.Inn != userInnRaw {
		t.Errorf("expected: %v, act: %v", userInnRaw, *act.Inn)
	}

	if !reflect.DeepEqual(act.Photo, userPhotoRaw) {
		t.Errorf("expected: %v, act: %v", userPhotoRaw, act.Photo)
	}

	if *act.CardNumber != userCardNumber {
		t.Errorf("expected: %v, act: %v", userCardNumber, *act.CardNumber)
	}

	if *act.CardExpires != userCardExpiresStamp {
		t.Errorf("expected: %v, act: %v", userCardExpiresStamp, *act.CardExpires)
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

	if !reflect.DeepEqual(act.Tolerances, userToleranceSnapshots) {
		t.Errorf("expected: %v, act: %v", userToleranceSnapshots, act.Tolerances)
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

func TestLoadFromSnapshot_Success(t *testing.T) {
	expectedUser, _ := NewUser(
		userID,
		userName,
		userPhone,
		&userRegistrationDate,
		userRating,
		userState,
		userIsRemoved,
		userVersion)

	expectedUser.ChangeName(userName)
	expectedUser.ChangePhone(userPhone)
	expectedUser.ChangePassport(passportEntity)
	expectedUser.ChangeDrivingLicense(drivingLicenseEntity)
	expectedUser.ChangeSnils(userSnils)
	expectedUser.ChangeInn(userInn)
	expectedUser.ChangePhoto(&userPhoto)
	expectedUser.ChangeCard(userCard)
	expectedUser.Activate()

	for i := 0; i < userPositiveRating; i++ {
		expectedUser.IncreaseRating()
	}

	for i := 0; i < userNegativeRating; i++ {
		expectedUser.DecreaseRating()
	}

	for _, val := range userAccessors {
		expectedUser.GrantAccess(*userOtherID, val)
	}

	snapshot := NewSnapshot(
		userID.String(),
		userFirstName,
		userMiddleName,
		userLastName,
		userCountryCode,
		userPhoneNumber,
		&passportSnapshot,
		&drivingLicenseSnapshot,
		&userSnilsRaw,
		&userInnRaw,
		userPhotoRaw,
		&userCardNumber,
		&userCardExpiresStamp,
		userRegistrationDateStamp,
		userPositiveRating,
		userNegativeRating,
		userToleranceSnapshots,
		userStateRaw,
		userIsRemoved,
		userVersion,
		userSnapshotDateStamp,
	)

	user, err := LoadFromSnapshot(&snapshot)
	if err != nil {
		t.Errorf("expected not error, act %v", err)
	}

	if act := user.GetID(); act != *userID {
		t.Errorf("expected: %v, act: %v", *userID, act)
	}

	if !reflect.DeepEqual(expectedUser, user) {
		t.Errorf("expected: %v, act: %v", expectedUser, user)
	}
}

func TestLoadFromSnapshot_Nil(t *testing.T) {
	act, err := LoadFromSnapshot(nil)
	if err != nil {
		t.Errorf("expected not error, act %v", err)
	}

	if act != nil {
		t.Errorf("expected: %v, act: %v", nil, act)
	}
}

func TestGetSnapshot_Success(t *testing.T) {
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

	for i := 0; i < userPositiveRating; i++ {
		user.IncreaseRating()
	}

	for i := 0; i < userNegativeRating; i++ {
		user.DecreaseRating()
	}

	for _, val := range userAccessors {
		user.GrantAccess(*userOtherID, val)
	}

	act, _ := GetSnapshot(user, userSnapshotDate)

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

	if !reflect.DeepEqual(act.Passport, &passportSnapshot) {
		t.Errorf("expected: %v, act: %v", &passportSnapshot, act.Passport)
	}

	if !reflect.DeepEqual(act.DrivingLicense, &drivingLicenseSnapshot) {
		t.Errorf("expected: %v, act: %v", &drivingLicenseSnapshot, act.DrivingLicense)
	}

	if *act.Snils != userSnilsRaw {
		t.Errorf("expected: %v, act: %v", userSnilsRaw, *act.Snils)
	}

	if *act.Inn != userInnRaw {
		t.Errorf("expected: %v, act: %v", userInnRaw, *act.Inn)
	}

	if !reflect.DeepEqual(act.Photo, userPhotoRaw) {
		t.Errorf("expected: %v, act: %v", userPhotoRaw, act.Photo)
	}

	if *act.CardNumber != userCardNumber {
		t.Errorf("expected: %v, act: %v", userCardNumber, *act.CardNumber)
	}

	if *act.CardExpires != userCardExpiresStamp {
		t.Errorf("expected: %v, act: %v", userCardExpiresStamp, *act.CardExpires)
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

	if !reflect.DeepEqual(act.Tolerances, userToleranceSnapshots) {
		t.Errorf("expected: %v, act: %v", userToleranceSnapshots, act.Tolerances)
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

func TestGetSnapshot_Nil(t *testing.T) {
	snapshot, _ := GetSnapshot(nil, userSnapshotDate)

	if snapshot != nil {
		t.Errorf("expected: %v, act: %v", nil, snapshot)
	}
}
