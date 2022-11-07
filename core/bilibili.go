package core

import (
	"github.com/go-resty/resty/v2"
	"github.com/wnstar/bili-utils/api/x/dm"
	"github.com/wnstar/bili-utils/api/x/relation"
	"github.com/wnstar/bili-utils/api/x/webInterface"
)

type Bilibili struct {
	client       *resty.Client
	Dm           dm.Dm
	WebInterface webInterface.WebInterface
	Relation     relation.Relation
}

func NewBilibili() *Bilibili {
	client := resty.New()
	client.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	client.SetHeader("referer", "https://space.bilibili.com")
	bilibili := &Bilibili{
		client:       client,
		Dm:           dm.Dm{Req: client},
		WebInterface: webInterface.WebInterface{Req: client},
		Relation:     relation.Relation{Req: client},
	}
	return bilibili
}
