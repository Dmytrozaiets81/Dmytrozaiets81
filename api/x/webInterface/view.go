package webInterface

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/wnstar/bili-utils/api/x/webType"
)

type WebInterface struct {
	Req *resty.Client
}

// View 获取视频详情
func (w WebInterface) View(id string) (*ViewResponse, error) {
	var response webType.WebResponse[ViewResponse]
	url := "http://api.bilibili.com/x/web-interface/view"
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
	resp, err := w.Req.R().SetQueryParams(map[string]string{key: id}).Get(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body(), &response)
	return &response.Data, err
}

func (w WebInterface) Card(mid string) (*CardResponse, error) {
	url := "http://api.bilibili.com/x/web-interface/card"
	resp, err := w.Req.R().
		SetQueryParams(map[string]string{"mid": mid, "photo": "false"}).
		Get(url)
	if err != nil {
		return nil, err
	}
	var data webType.WebResponse[CardResponse]
	err = json.Unmarshal(resp.Body(), &data)
	return &data.Data, err
}

func (w WebInterface) ArchiveStat(bvid string, aid int) (*ArchiveStatResponse, error) {
	url := "http://api.bilibili.com/x/web-interface/archive/stat"
	client := resty.New()
	params := map[string]string{}
	if aid != 0 {
		params["aid"] = fmt.Sprintf("%v", aid)
	} else {
		params["bvid"] = bvid
	}
	resp, err := client.R().
		SetQueryParams(params).
		Get(url)
	if err != nil {
		return nil, err
	}
	var data webType.WebResponse[ArchiveStatResponse]
	err = json.Unmarshal(resp.Body(), &data)
	return &data.Data, err
}

func (w WebInterface) Online() (*OnlineResponse, error) {
	url := "http://api.bilibili.com/x/web-interface/online"
	resp, err := w.Req.R().Get(url)
	if err != nil {
		return nil, err
	}
	var data webType.WebResponse[OnlineResponse]
	err = json.Unmarshal(resp.Body(), &data)
	return &data.Data, err
}

func (w WebInterface) DynamicRegion(rid string, page int, pagesize int) (*DynamicRegionResponse, error) {
	url := "http://api.bilibili.com/x/web-interface/dynamic/region"
	params := map[string]string{"rid": rid}
	if page == 0 {
		params["pn"] = "1"
	} else {
		params["pn"] = fmt.Sprintf("%v", page)
	}
	if pagesize == 0 {
		params["ps"] = "5"
	} else {
		params["ps"] = fmt.Sprintf("%v", pagesize)
	}
	resp, err := w.Req.R().Get(url)
	if err != nil {
		return nil, err
	}
	var data webType.WebResponse[DynamicRegionResponse]
	err = json.Unmarshal(resp.Body(), &data)
	return &data.Data, err
}

func (w WebInterface) NewList(rid string, pagesize, page int) (*NewListResponse, error) {
	url := "https://api.bilibili.com/x/web-interface/newlist"
	params := map[string]string{"rid": rid, "type": "0", "ps": fmt.Sprintf("%v", pagesize), "pn": fmt.Sprintf("%v", page)}
	resp, err := w.Req.R().SetQueryParams(params).Get(url)
	if err != nil {
		return nil, err
	}
	var data webType.WebResponse[NewListResponse]
	err = json.Unmarshal(resp.Body(), &data)
	return &data.Data, err
}
