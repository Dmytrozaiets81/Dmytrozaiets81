package core

import (
	"github.com/wnstar/bili-utils/api/x/dm"
	"github.com/wnstar/bili-utils/api/x/relation"
	"github.com/wnstar/bili-utils/api/x/webInterface"
	"net/http"
)

type Bilibili struct {
	client       *http.Client
	Dm           dm.Dm
	WebInterface webInterface.WebInterface
	Relation     relation.Relation
}

func NewBilibili() *Bilibili {
	bilibili := &Bilibili{
		client:       http.DefaultClient,
		Dm:           dm.Dm{},
		WebInterface: webInterface.WebInterface{},
		Relation:     relation.Relation{},
	}
	return bilibili
}
