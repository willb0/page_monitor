package main

import (
	"context"
	"log"
	"page_monitor/page_refresh"
	"time"

	"github.com/itzmeanjan/pub0sub/publisher"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	pub, err := publisher.New(ctx, "tcp", "localhost:13000")
	if err != nil {
		log.Println("what da hell")
		log.Println(err)
		return
	}
	refresher := page_refresh.NewPageRefresher("http://localhost:3001")
	refresher.WatchForChangesAndNotify(pub,2)
	cancel()
	<-time.After(time.Second)
}
