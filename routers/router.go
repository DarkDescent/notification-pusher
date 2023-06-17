package routers

import (
	"notifications-pusher/controllers"

	"notifications-pusher/pkg/settings"

	"database/sql"
	"notifications-pusher/db/models"

	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	db, err := sql.Open("postgres", settings.ServiceConfig.Database.ConnectionString)
	if err != nil {
		processError(err)
	}
	databaseConn := models.New(db)

	messageController := controllers.MessageController{ServiceKey: settings.ServiceConfig.Service.FirebaseCredentials, Database: databaseConn}
	tokenController := controllers.TokenController{Database: databaseConn}
	r.POST("/messages", messageController.SendMessage)
	r.POST("/tokens", tokenController.SaveToken)

	expiringTask := Expiring{Database: databaseConn}
	expiringTask.InitExpiring()

	return r
}

func processError(err error) {
	log.Fatal(err)
	os.Exit(2)
}
