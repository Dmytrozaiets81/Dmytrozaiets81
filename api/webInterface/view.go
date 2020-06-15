package webInterface

import (
	"github.com/wnstar/gorequests"
	"strings"
)

type ViewResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Bvid string `json:""`
		Aid  int64  `json:"aid"`
	}
	Pages []struct {
		Cid int64 `json:"cid"`
	}
}

func NewViewFromId(id string) (*ViewResponse, error) {
	var response ViewResponse
	key := ""
	switch strings.ToLower(id[:2]) {
	case "bv":
		key = "bvid"
	case "av":
		key = "aid"
		id = id[2:]
	default:
		key = "aid"
	}

	resp, err := gorequests.Get("https://api.bilibili.com/x/web-interface/view").AddParams(key, id).Start()
	if err != nil {
		return nil, err
	}
	resp.Json(&response)
	return &response, nil
}
