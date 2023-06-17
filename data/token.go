package data

import (
	"notifications-pusher/db/models"
	"time"

	"github.com/google/uuid"
)

type Token struct {
	Cid           uuid.UUID `json:"cid" binding:"required"`
	FirebaseToken string    `json:"token" binding:"required"`
}

func CreateDatabaseToken(cid uuid.UUID, firebaseToken string) models.CreateFirebaseTokenParams {
	return models.CreateFirebaseTokenParams{
		Token:  firebaseToken,
		Cid:    cid,
		Active: true,
		// TODO: expires at must be set as parameter
		Expiresat: time.Now().AddDate(0, 0, 30),
	}
}

func DuplicateDatabaseTokenParams(cid uuid.UUID, firebaseToken string) models.FindDuplicatedTokenParams {
	return models.FindDuplicatedTokenParams{
		Token: firebaseToken,
		Cid:   cid,
	}
}

func ActivityParams(id int64, active bool) models.SetActivitityParams {
	return models.SetActivitityParams{
		Active: active,
		ID:     id,
	}
}
