package controllers

import (
	"context"
	"net/http"
	"notifications-pusher/data"

	"notifications-pusher/db/models"

	"fmt"

	"database/sql"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TokenController struct {
	Database *models.Queries
}

func (tc *TokenController) SaveToken(c *gin.Context) {
	var message data.Token
	if err := c.BindJSON(&message); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkDuplicateParams := data.DuplicateDatabaseTokenParams(message.Cid, message.FirebaseToken)
	found, err := tc.Database.FindDuplicatedToken(context.Background(), checkDuplicateParams)

	if err != nil && err != sql.ErrNoRows {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if found.ID != 0 {
		log.Info(fmt.Sprintf("Found already existing message by token %s and cid %s", message.FirebaseToken, message.Cid))
		c.Status(http.StatusAccepted)
		return
	}

	toSave := data.CreateDatabaseToken(message.Cid, message.FirebaseToken)

	newToken, err := tc.Database.CreateFirebaseToken(context.Background(), toSave)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": newToken.Token})
}
