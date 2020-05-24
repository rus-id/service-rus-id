package main

import (
	"time"

	"github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport"
	passValueTypes "github.com/bgoldovsky/service-rus-id/internal/domain/entities/passport/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/domain/user"
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/logger"
	"github.com/sirupsen/logrus"
)

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
	const isRemoved = true
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
	u.Activate()

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

	passID, err := passValueTypes.NewPassportID("2233", "123456")
	if err != nil {
		logError("passID", passID, err)
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

	passValidation := passValueTypes.NewPassportValidation(true, false, true, false)
	pass, err := passport.NewPassport(passID, name, &birthday, passIssue, registration, passValidation)
	if err != nil {
		logError("pass", pass, err)
	}

	u.ChangePassport(pass)
	logger.Log.WithFields(logrus.Fields{"User": u}).Info("user updated")

	snapshot, err := user.GetSnapshot(u, time.Now().UTC())
	if err != nil {
		logError("snapshot", snapshot, err)
	}

	logger.Log.WithFields(logrus.Fields{"Snapshot": snapshot}).Warn("snapshot retrieved")
}

func logError(key string, val interface{}, err error) {
	logger.Log.WithField(key, val).WithError(err).Errorln("user not created")
}
