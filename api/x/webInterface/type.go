package webInterface

type ViewResponse struct {
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
}

type CardResponse struct {
	Card struct {
		Mid         string        `json:"mid"`
		Name        string        `json:"name"`
		Approve     bool          `json:"approve"`
		Sex         string        `json:"sex"`
		Rank        string        `json:"rank"`
		Face        string        `json:"face"`
		FaceNft     int           `json:"face_nft"`
		FaceNftType int           `json:"face_nft_type"`
		DisplayRank string        `json:"DisplayRank"`
		Regtime     int           `json:"regtime"`
		Spacesta    int           `json:"spacesta"`
		Birthday    string        `json:"birthday"`
		Place       string        `json:"place"`
		Description string        `json:"description"`
		Article     int           `json:"article"`
		Attentions  []interface{} `json:"attentions"`
		Fans        int           `json:"fans"`
		Friend      int           `json:"friend"`
		Attention   int           `json:"attention"`
		Sign        string        `json:"sign"`
		LevelInfo   struct {
			CurrentLevel int `json:"current_level"`
			CurrentMin   int `json:"current_min"`
			CurrentExp   int `json:"current_exp"`
			NextExp      int `json:"next_exp"`
		} `json:"level_info"`
		Pendant struct {
			Pid               int    `json:"pid"`
			Name              string `json:"name"`
			Image             string `json:"image"`
			Expire            int    `json:"expire"`
			ImageEnhance      string `json:"image_enhance"`
			ImageEnhanceFrame string `json:"image_enhance_frame"`
		} `json:"pendant"`
		Nameplate struct {
			Nid        int    `json:"nid"`
			Name       string `json:"name"`
			Image      string `json:"image"`
			ImageSmall string `json:"image_small"`
			Level      string `json:"level"`
			Condition  string `json:"condition"`
		} `json:"nameplate"`
		Official struct {
			Role  int    `json:"role"`
			Title string `json:"title"`
			Desc  string `json:"desc"`
			Type  int    `json:"type"`
		} `json:"Official"`
		OfficialVerify struct {
			Type int    `json:"type"`
			Desc string `json:"desc"`
		} `json:"official_verify"`
		Vip struct {
			Type       int   `json:"type"`
			Status     int   `json:"status"`
			DueDate    int64 `json:"due_date"`
			VipPayType int   `json:"vip_pay_type"`
			ThemeType  int   `json:"theme_type"`
			Label      struct {
				Path                  string `json:"path"`
				Text                  string `json:"text"`
				LabelTheme            string `json:"label_theme"`
				TextColor             string `json:"text_color"`
				BgStyle               int    `json:"bg_style"`
				BgColor               string `json:"bg_color"`
				BorderColor           string `json:"border_color"`
				UseImgLabel           bool   `json:"use_img_label"`
				ImgLabelURIHans       string `json:"img_label_uri_hans"`
				ImgLabelURIHant       string `json:"img_label_uri_hant"`
				ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
				ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
			} `json:"label"`
			AvatarSubscript    int    `json:"avatar_subscript"`
			NicknameColor      string `json:"nickname_color"`
			Role               int    `json:"role"`
			AvatarSubscriptURL string `json:"avatar_subscript_url"`
			TvVipStatus        int    `json:"tv_vip_status"`
			TvVipPayType       int    `json:"tv_vip_pay_type"`
			VipType            int    `json:"vipType"`
			VipStatus          int    `json:"vipStatus"`
		} `json:"vip"`
		IsSeniorMember int `json:"is_senior_member"`
	} `json:"card"`
	Following    bool `json:"following"`
	ArchiveCount int  `json:"archive_count"`
	ArticleCount int  `json:"article_count"`
	Follower     int  `json:"follower"`
	LikeNum      int  `json:"like_num"`
}

type ArchiveStatResponse struct {
	Aid        int    `json:"aid"`
	Bvid       string `json:"bvid"`
	View       int    `json:"view"`
	Danmaku    int    `json:"danmaku"`
	Reply      int    `json:"reply"`
	Favorite   int    `json:"favorite"`
	Coin       int    `json:"coin"`
	Share      int    `json:"share"`
	Like       int    `json:"like"`
	NowRank    int    `json:"now_rank"`
	HisRank    int    `json:"his_rank"`
	NoReprint  int    `json:"no_reprint"`
	Copyright  int    `json:"copyright"`
	ArgueMsg   string `json:"argue_msg"`
	Evaluation string `json:"evaluation"`
}

type OnlineResponse struct {
	RegionCount map[string]int `json:"region_count"`
}

type DynamicRegionResponse struct {
	Page struct {
		Num   int `json:"num"`
		Size  int `json:"size"`
		Count int `json:"count"`
	} `json:"page"`
	Archives []struct {
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		State     int    `json:"state"`
		Duration  int    `json:"duration"`
		MissionID int    `json:"mission_id,omitempty"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			ArcPay        int `json:"arc_pay"`
			PayFreeWatch  int `json:"pay_free_watch"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid      int `json:"aid"`
			View     int `json:"view"`
			Danmaku  int `json:"danmaku"`
			Reply    int `json:"reply"`
			Favorite int `json:"favorite"`
			Coin     int `json:"coin"`
			Share    int `json:"share"`
			NowRank  int `json:"now_rank"`
			HisRank  int `json:"his_rank"`
			Like     int `json:"like"`
			Dislike  int `json:"dislike"`
		} `json:"stat"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		SeasonID    int         `json:"season_id,omitempty"`
		ShortLink   string      `json:"short_link"`
		ShortLinkV2 string      `json:"short_link_v2"`
		FirstFrame  string      `json:"first_frame"`
		PubLocation string      `json:"pub_location"`
		Bvid        string      `json:"bvid"`
		SeasonType  int         `json:"season_type"`
		IsOgv       bool        `json:"is_ogv"`
		OgvInfo     interface{} `json:"ogv_info"`
		RcmdReason  string      `json:"rcmd_reason"`
	} `json:"archives"`
}

type NewListResponse struct {
	Archives []struct {
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		State     int    `json:"state"`
		Duration  int    `json:"duration"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			ArcPay        int `json:"arc_pay"`
			PayFreeWatch  int `json:"pay_free_watch"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid      int `json:"aid"`
			View     int `json:"view"`
			Danmaku  int `json:"danmaku"`
			Reply    int `json:"reply"`
			Favorite int `json:"favorite"`
			Coin     int `json:"coin"`
			Share    int `json:"share"`
			NowRank  int `json:"now_rank"`
			HisRank  int `json:"his_rank"`
			Like     int `json:"like"`
			Dislike  int `json:"dislike"`
		} `json:"stat"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		ShortLink   string      `json:"short_link"`
		ShortLinkV2 string      `json:"short_link_v2"`
		FirstFrame  string      `json:"first_frame"`
		PubLocation string      `json:"pub_location"`
		Bvid        string      `json:"bvid"`
		SeasonType  int         `json:"season_type"`
		IsOgv       bool        `json:"is_ogv"`
		OgvInfo     interface{} `json:"ogv_info"`
		RcmdReason  string      `json:"rcmd_reason"`
		UpFromV2    int         `json:"up_from_v2,omitempty"`
		MissionID   int         `json:"mission_id,omitempty"`
		SeasonID    int         `json:"season_id,omitempty"`
	} `json:"archives"`
	Page struct {
		Count int `json:"count"`
		Num   int `json:"num"`
		Size  int `json:"size"`
	} `json:"page"`
}
