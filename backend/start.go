package backend

import (
	"time"

	"gitlab.com/gaming0skar123/go/pingbot/config"
)

func Start() {
	// Ping on Start
	ping()

	ticker := time.NewTicker(config.PingBot_Ticker)

	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			ping()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}