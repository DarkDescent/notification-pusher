package routers

import (
	"context"
	"notifications-pusher/data"
	"notifications-pusher/db/models"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	checkInterval = 1 * time.Minute // How often to check for expired tokens
)

var (
	ticker *time.Ticker
	done   chan bool
)

type Expiring struct {
	Database *models.Queries
}

func (e *Expiring) InitExpiring() {
	go e.startChecking()
}

// For now we do not need this function, we set our goroutine to work while our server works
// func StopExpiring() {
//     done <- true
// }

func (e *Expiring) startChecking() {
	ticker = time.NewTicker(checkInterval)
	done = make(chan bool)

	for {
		select {
		case <-ticker.C:
			e.checkExpiredTokens()
		case <-done:
			ticker.Stop()
			return
		}
	}
}

func (e *Expiring) checkExpiredTokens() {
	tokens, err := e.Database.FindExpiredToken(context.Background(), time.Now())
	if err != nil {
		log.Error(err)
		return
	}

	for _, t := range tokens {
		_, err = e.Database.SetActivitity(context.Background(), data.ActivityParams(t.ID, false))
		if err != nil {
			log.Error(err)
			continue
		}
	}
}
