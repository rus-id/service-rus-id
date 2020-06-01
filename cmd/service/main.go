package main

import (
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license"
	dlValueTypes "github.com/bgoldovsky/service-rus-id/internal/domain/entities/driving_license/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"
	passValueTypes "github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/logger"
	"github.com/bgoldovsky/service-rus-id/internal/repository/in_memory"
)

// TODO: Почистить помойку
func main() {
	userID := valuetypes.CreateUserID()
	name, err := valuetypes.NewName("Boris", nil, "Goldovsky")
	if err != nil {
		logError("name", name, err)
	}

	phone, err := valuetypes.NewPhone(valuetypes.RusCountryCode, "9039615395")
	if err != nil {
		logError("phone", phone, err)
	}

	registrationDate := time.Now()
	state := valuetypes.UserStateBlocked
	const isRemoved = false
	version := int64(777)

	rating, err := valuetypes.NewRating(5, 2)
	if err != nil {
		logError("rating", rating, err)
	}

	u, err := user.NewUser(
		userID,
		name,
		phone,
		&registrationDate,
		rating,
		state,
		isRemoved,
		version)

	if err != nil {
		logError("user", u, err)
	}

	snils, err := valuetypes.NewSnils("59650418527")
	if err != nil {
		logError("snils", snils, err)
	}
	u.ChangeSnils(snils)

	inn, err := valuetypes.NewInn("889373498613")
	if err != nil {
		logError("inn", inn, err)
	}
	u.ChangeInn(inn)

	pass := getPassport()
	u.ChangePassport(pass)

	dl := getDrivingLicense()
	u.ChangeDrivingLicense(dl)
	u.GrantFullAccess(*valuetypes.CreateUserID())

	expires := time.Date(2020, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	card, err := valuetypes.NewCard("4444333322221111", expires)
	u.ChangeCard(card)

	photo := valuetypes.Photo{1, 2, 3}
	u.ChangePhoto(&photo)

	store := make(map[valuetypes.UserID]*user.Snapshot)
	repo, err := in_memory.NewInMemoryRepository(store)
	if err != nil {
		logError("user", u, err)
	}

	err = repo.Save(u)
	if err != nil {
		logError("user", u, err)
	}

	isExist, err := repo.IsExist(u.GetID())
	if err != nil {
		logError("user", u, err)
	}

	logger.Log.Infof("User exist %v", isExist)

	loaded, err := repo.Find(u.GetID())
	if err != nil {
		logError("id", u.GetID(), err)
	}

	logger.Log.Info(loaded.String())
}

func logError(key string, val interface{}, err error) {
	logger.Log.WithField(key, val).WithError(err).Errorln("user not created")
}

func getPassport() *passport.Passport {
	id, err := passValueTypes.NewPassportID("2233", "123456")
	if err != nil {
		logError("id", id, err)
	}

	name, err := valuetypes.NewName("Boris", nil, "Goldovsky")
	if err != nil {
		logError("name", name, err)
	}

	birthday := time.Date(1986, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	issueDate := time.Date(2010, time.Month(6), 19, 1, 10, 30, 0, time.UTC)
	passIssue, err := passValueTypes.NewPassportIssue("MVD", issueDate, "770-21")
	if err != nil {
		logError("passIssue", passIssue, err)
	}

	registration, err := valuetypes.NewAddress("Shipilovskaya st.")
	if err != nil {
		logError("registration", registration, err)
	}

	validation := passValueTypes.NewPassportValidation(true, false, true, false)

	pass, err := passport.NewPassport(id, name, &birthday, passIssue, registration, validation)
	if err != nil {
		logError("pass", pass, err)
	}

	return pass
}

func getDrivingLicense() *driving_license.DrivingLicense {
	id, err := dlValueTypes.NewDrivingLicenseID("2233", "123456")
	if err != nil {
		logError("id", id, err)
	}

	name, err := valuetypes.NewName("Boris", nil, "Goldovsky")
	if err != nil {
		logError("name", name, err)
	}

	residence, err := dlValueTypes.NewResidence("Russia")
	if err != nil {
		logError("residence", residence, err)
	}

	category := dlValueTypes.DrivingLicenseA
	birthday := time.Date(1986, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	issued := time.Date(2010, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	expires := time.Date(2025, time.Month(4), 9, 1, 10, 30, 0, time.UTC)
	specialMarks := "empty mark"
	validation := dlValueTypes.NewDrivingLicenseValidation(true, false)

	dl, err := driving_license.NewDrivingLicense(
		id,
		category,
		name,
		&birthday,
		&issued,
		&expires,
		residence,
		specialMarks,
		validation)

	if err != nil {
		logError("residence", residence, err)
	}

	return dl
}
