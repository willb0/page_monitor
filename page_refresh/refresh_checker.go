package page_refresh

import (
	"context"
	"crypto/sha256"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"

	"github.com/redis/go-redis/v9"
)

type PageRefresher struct {
	page_url        string
	current_content []byte
}

func NewPageRefresher(page_url string) *PageRefresher {
	return &PageRefresher{
		page_url:        page_url,
		current_content: []byte(page_url),
	}
}

func (p *PageRefresher) CheckForChanges() bool {
	html := p.GetHTML()
	hasher := sha256.New()
	hasher.Write([]byte(html))
	hashed_html := hasher.Sum(nil)
	res := reflect.DeepEqual(hashed_html, p.current_content)
	if !res {
		p.current_content = hashed_html
		return true
	}
	return false
}

func (p *PageRefresher) GetHTML() []byte {
	resp, err := http.Get(p.page_url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return html
}

func (p *PageRefresher) WatchForChangesAndNotify(ctx context.Context,r *redis.Client,refresh_rate int){
	for {
		time.Sleep(time.Second * time.Duration(refresh_rate))
		println("checking for changes")
		if p.CheckForChanges() {
			println("it changed!")
			r.Publish(ctx,"page_refresh",p.GetHTML())
		}
	}
} 
