package relation

import (
	"encoding/json"
	"fmt"
	"github.com/wnstar/bili-utils/request"
)

type Relation struct {
}

type FollowersResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		List []struct {
			Mid            int         `json:"mid"`
			Attribute      int         `json:"attribute"`
			Mtime          int         `json:"mtime"`
			Tag            interface{} `json:"tag"`
			Special        int         `json:"special"`
			Uname          string      `json:"uname"`
			Face           string      `json:"face"`
			Sign           string      `json:"sign"`
			OfficialVerify struct {
				Type int    `json:"type"`
				Desc string `json:"desc"`
			} `json:"official_verify"`
			Vip struct {
				VipType       int    `json:"vipType"`
				VipDueDate    int64  `json:"vipDueDate"`
				DueRemark     string `json:"dueRemark"`
				AccessStatus  int    `json:"accessStatus"`
				VipStatus     int    `json:"vipStatus"`
				VipStatusWarn string `json:"vipStatusWarn"`
				ThemeType     int    `json:"themeType"`
				Label         struct {
					Path string `json:"path"`
				} `json:"label"`
			} `json:"vip"`
		} `json:"list"`
		ReVersion int64 `json:"re_version"`
		Total     int   `json:"total"`
	} `json:"data"`
}

func (Relation) Followers(uid string, page int, limit int) (*FollowersResponse, error) {
	var tmp FollowersResponse
	if limit == 0 {
		limit = 20
	}
	path := fmt.Sprintf("/x/relation/followers?vmid=%s&pn=%d&ps=%d&order=desc&jsonp=jsonp", uid, page, limit)
	body, _, err := request.Get(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &tmp)
	return &tmp, err
}
