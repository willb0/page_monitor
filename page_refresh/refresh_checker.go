package page_refresh

import (
	"crypto/sha256"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"

	"github.com/itzmeanjan/pub0sub/ops"
	"github.com/itzmeanjan/pub0sub/publisher"
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
	resp, err := http.Get(p.page_url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
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

func (p *PageRefresher) WatchForChangesAndNotify(pub *publisher.Publisher,refresh_rate int){
	for {
		time.Sleep(time.Second * time.Duration(refresh_rate))
		println("checking for changes")
		if p.CheckForChanges() {
			println("it changed!")
			data := []byte("the page refreshed!!!")
			topics := []string{"page_refresh"}
			msg := ops.Msg{Topics: topics, Data: data}
			pub.Publish(&msg)
		}
	}
} 