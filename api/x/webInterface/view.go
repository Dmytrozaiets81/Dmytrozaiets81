package webInterface

import (
	"encoding/json"
	"fmt"
	"github.com/wnstar/bili-utils/request"
	"strings"
)

type WebInterface struct {
}

type ViewResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Bvid      string `json:"bvid"`      //视频bvid
		Aid       int    `json:"aid"`       //视频avid
		Videos    int    `json:"videos"`    //分p数
		Tid       int    `json:"tid"`       //分区id
		Tname     string `json:"tname"`     //子分区
		Copyright int    `json:"copyright"` //是否自制 1自制
		Pic       string `json:"pic"`       //封面
		Title     string `json:"title"`     //标题
		Pubdate   int    `json:"pubdate"`   //上传时间戳
		Ctime     int    `json:"ctime"`     //通过审核时间
		Desc      string `json:"desc"`      // 描述
		State     int    `json:"state"`
		Attribute int    `json:"attribute"`
		Duration  int    `json:"duration"` //时长（秒
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"` //是否允许下载
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
		} `json:"rights"`
		Owner struct { //发布者信息
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct { //视频一些信息
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
		} `json:"stat"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		NoCache bool `json:"no_cache"`
		Pages   []struct {
			Cid       int    `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool          `json:"allow_submit"`
			List        []interface{} `json:"list"`
		} `json:"subtitle"`
	} `json:"data"`
}

//View 获取视频详情
func (w WebInterface) View(id string) (*ViewResponse, error) {
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
	body, _, err := request.Get(fmt.Sprintf("/x/web-interface/view?%s=%s", key, id))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	return &response, err
}
