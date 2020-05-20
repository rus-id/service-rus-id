package main

import (
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/logger"
	"github.com/sirupsen/logrus"
)

func main() {

	//userID := valuetypes.CreateUserID()
	_, err := valuetypes.NewSnils("59650418527")
	if err != nil {
		logger.Log.
			WithFields(logrus.Fields{
				"SNILS": err,
			}).Errorln("user not created")
		return
	}

	//user, err := aggregates.NewUser(userID, snils)
	//user := domain.NewUser(userID, snils)

	logger.Log.
		WithFields(logrus.Fields{
			//	"User": user,
		}).Info("user created")

	logger.Log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Warn("A group of walrus emerges from the ocean")
}
