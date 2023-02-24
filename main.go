package main

import (
	"context"
	"flag"
	"page_monitor/page_refresh"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	url := flag.String("url", "http://localhost:8000","url to monitor")
	rate := flag.Int("refresh_rate",5,"seconds to wait before checking page again")
	ctx, cancel := context.WithCancel(context.Background())
	rdb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})
	flag.Parse()
	refresher := page_refresh.NewPageRefresher(*url)
	refresher.WatchForChangesAndNotify(ctx,rdb,*rate)
	cancel()
	<-time.After(time.Second)

}
