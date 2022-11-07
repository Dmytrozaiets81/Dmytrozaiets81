package relation

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/wnstar/bili-utils/api/x/webType"
)

type Relation struct {
	Req *resty.Client
}

type FollowersResponse struct {
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
}

func (r Relation) Followers(uid string, page int, limit int) (*FollowersResponse, error) {
	var tmp webType.WebResponse[FollowersResponse]
	if limit == 0 {
		limit = 20
	}
	url := "http://api.bilibili.com/x/relation/followers"
	resp, err := r.Req.R().SetQueryParams(map[string]string{"vmid": uid, "pn": fmt.Sprintf("%v", page), "ps": fmt.Sprintf("%v", limit), "order": "desc", "jsonp": "jsonp"}).Get(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body(), &tmp)
	return &tmp.Data, err
}
