package controllers

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"notifications-pusher/data"
	"time"

	"notifications-pusher/db/models"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

type MessageController struct {
	ServiceKey string
	Database   *models.Queries
}

func (mc *MessageController) SendMessage(c *gin.Context) {
	decodedKey, err := base64.StdEncoding.DecodeString(mc.ServiceKey)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - check logs."})
	}
	opt := option.WithCredentialsJSON(decodedKey)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - check logs."})
	}
	fcmClient, err := app.Messaging(context.Background())
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - check logs."})
	}

	var message data.Message
	if err := c.BindJSON(&message); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	firebaseToken, err := mc.Database.FindTokenByCid(context.Background(), message.Cid)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	if now.After(firebaseToken.Expiresat) {
		log.Error(fmt.Sprintf("We have got request to send message to device with expired token - expired cid - %s, expired token - %s", firebaseToken.Cid, firebaseToken.Token))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired. You need to restart your device"})
		return
	}

	_, err = fcmClient.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: message.Title,
			Body:  message.Body,
		},
		Token: firebaseToken.Token,
	})

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
