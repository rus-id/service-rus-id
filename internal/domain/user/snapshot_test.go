package user_test

import (
	"reflect"
	"testing"

	"github.com/bgoldovsky/service-rus-id/internal/domain/mock"

	. "github.com/bgoldovsky/service-rus-id/internal/domain/user"
)

func TestNewSnapshot(t *testing.T) {
	act := NewSnapshot(
		mock.UserIDRaw,
		mock.UserFirstName,
		mock.UserMiddleName,
		mock.UserLastName,
		mock.UserCountryCode,
		mock.UserPhoneNumber,
		&mock.PassportSnapshot,
		&mock.DrivingLicenseSnapshot,
		&mock.UserSnilsRaw,
		&mock.UserInnRaw,
		mock.UserPhotoRaw,
		&mock.UserCardNumber,
		&mock.UserCardExpiresStamp,
		mock.UserRegistrationDateStamp,
		mock.UserPositiveRating,
		mock.UserNegativeRating,
		mock.UserToleranceSnapshots,
		mock.UserStateRaw,
		mock.UserIsRemoved,
		mock.UserVersion,
		mock.UserSnapshotDateStamp,
	)

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

	if !reflect.DeepEqual(act.Passport, &mock.PassportSnapshot) {
		t.Errorf("expected: %v, act: %v", &mock.PassportSnapshot, act.Passport)
	}

	if !reflect.DeepEqual(act.DrivingLicense, &mock.DrivingLicenseSnapshot) {
		t.Errorf("expected: %v, act: %v", &mock.DrivingLicenseSnapshot, act.DrivingLicense)
	}

	if *act.Snils != mock.UserSnilsRaw {
		t.Errorf("expected: %v, act: %v", mock.UserSnilsRaw, *act.Snils)
	}

	if *act.Inn != mock.UserInnRaw {
		t.Errorf("expected: %v, act: %v", mock.UserInnRaw, *act.Inn)
	}

	if !reflect.DeepEqual(act.Photo, mock.UserPhotoRaw) {
		t.Errorf("expected: %v, act: %v", mock.UserPhotoRaw, act.Photo)
	}

	if *act.CardNumber != mock.UserCardNumber {
		t.Errorf("expected: %v, act: %v", mock.UserCardNumber, *act.CardNumber)
	}

	if *act.CardExpires != mock.UserCardExpiresStamp {
		t.Errorf("expected: %v, act: %v", mock.UserCardExpiresStamp, *act.CardExpires)
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

	if !reflect.DeepEqual(act.Tolerances, mock.UserToleranceSnapshots) {
		t.Errorf("expected: %v, act: %v", mock.UserToleranceSnapshots, act.Tolerances)
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

func TestLoadFromSnapshot_Success(t *testing.T) {
	expectedUser, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		mock.UserIsRemoved,
		mock.UserVersion)

	expectedUser.ChangeName(mock.UserName)
	expectedUser.ChangePhone(mock.UserPhone)
	expectedUser.ChangePassport(mock.PassportEntity)
	expectedUser.ChangeDrivingLicense(mock.DrivingLicenseEntity)
	expectedUser.ChangeSnils(mock.UserSnils)
	expectedUser.ChangeInn(mock.UserInn)
	expectedUser.ChangePhoto(&mock.UserPhoto)
	expectedUser.ChangeCard(mock.UserCard)
	expectedUser.Activate()

	for _, val := range mock.UserAccessors {
		expectedUser.GrantAccess(*mock.UserOtherID, val)
	}

	snapshot := NewSnapshot(
		mock.UserID.String(),
		mock.UserFirstName,
		mock.UserMiddleName,
		mock.UserLastName,
		mock.UserCountryCode,
		mock.UserPhoneNumber,
		&mock.PassportSnapshot,
		&mock.DrivingLicenseSnapshot,
		&mock.UserSnilsRaw,
		&mock.UserInnRaw,
		mock.UserPhotoRaw,
		&mock.UserCardNumber,
		&mock.UserCardExpiresStamp,
		mock.UserRegistrationDateStamp,
		mock.UserPositiveRating,
		mock.UserNegativeRating,
		mock.UserToleranceSnapshots,
		mock.UserStateRaw,
		mock.UserIsRemoved,
		mock.UserVersion,
		mock.UserSnapshotDateStamp,
	)

	user, err := LoadFromSnapshot(&snapshot)
	if err != nil {
		t.Errorf("expected not error, act %v", err)
	}

	if act := user.GetID(); act != *mock.UserID {
		t.Errorf("expected: %v, act: %v", *mock.UserID, act)
	}

	if !reflect.DeepEqual(expectedUser, user) {
		t.Errorf("expected: %v, act: %v", expectedUser, user)
	}
}

func TestLoadFromSnapshot_Nil(t *testing.T) {
	act, err := LoadFromSnapshot(nil)
	if err != ErrInvalidSnapshot {
		t.Errorf("expected error: %v, act %v", ErrInvalidSnapshot, err)
	}

	if act != nil {
		t.Errorf("expected: %v, act: %v", nil, act)
	}
}

func TestGetSnapshot_Success(t *testing.T) {
	user, _ := NewUser(
		mock.UserID,
		mock.UserName,
		mock.UserPhone,
		&mock.UserRegistrationDate,
		mock.UserRating,
		mock.UserState,
		mock.UserIsRemoved,
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

	for _, val := range mock.UserAccessors {
		user.GrantAccess(*mock.UserOtherID, val)
	}

	act, _ := GetSnapshot(user, mock.UserSnapshotDate)

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

	if !reflect.DeepEqual(act.Passport, &mock.PassportSnapshot) {
		t.Errorf("expected: %v, act: %v", &mock.PassportSnapshot, act.Passport)
	}

	if !reflect.DeepEqual(act.DrivingLicense, &mock.DrivingLicenseSnapshot) {
		t.Errorf("expected: %v, act: %v", &mock.DrivingLicenseSnapshot, act.DrivingLicense)
	}

	if *act.Snils != mock.UserSnilsRaw {
		t.Errorf("expected: %v, act: %v", mock.UserSnilsRaw, *act.Snils)
	}

	if *act.Inn != mock.UserInnRaw {
		t.Errorf("expected: %v, act: %v", mock.UserInnRaw, *act.Inn)
	}

	if !reflect.DeepEqual(act.Photo, mock.UserPhotoRaw) {
		t.Errorf("expected: %v, act: %v", mock.UserPhotoRaw, act.Photo)
	}

	if *act.CardNumber != mock.UserCardNumber {
		t.Errorf("expected: %v, act: %v", mock.UserCardNumber, *act.CardNumber)
	}

	if *act.CardExpires != mock.UserCardExpiresStamp {
		t.Errorf("expected: %v, act: %v", mock.UserCardExpiresStamp, *act.CardExpires)
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

	if !reflect.DeepEqual(act.Tolerances, mock.UserToleranceSnapshots) {
		t.Errorf("expected: %v, act: %v", mock.UserToleranceSnapshots, act.Tolerances)
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

func TestGetSnapshot_Nil(t *testing.T) {
	snapshot, _ := GetSnapshot(nil, mock.UserSnapshotDate)

	if snapshot != nil {
		t.Errorf("expected: %v, act: %v", nil, snapshot)
	}
}
