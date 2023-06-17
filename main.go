package main

import (
	"notifications-pusher/pkg/settings"
	"notifications-pusher/routers"

	log "github.com/sirupsen/logrus"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	settings.InitConfig()
}

func main() {
	gin.SetMode(settings.ServiceConfig.Server.RunMode)

	routersInit := routers.InitRouter()

	readTimeout := settings.ServiceConfig.Server.ReadTimeout
	writeTimeout := settings.ServiceConfig.Server.WriteTimeout
	endPoint := fmt.Sprintf(":%d", settings.ServiceConfig.Server.Port)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Info("start http server listening ", endPoint)

	server.ListenAndServe()
}
