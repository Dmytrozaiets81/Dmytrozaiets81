package dm

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/golang/protobuf/proto"
)

//go:generate protoc --go_out=./ bilibili.proto

type Dm struct {
	Req *resty.Client
}

func (d Dm) SegSo(cid string) ([]*DanmakuElem, error) {
	result := make([]*DanmakuElem, 0)
	i := 0
	for {
		var tmp DmSegMobileReply
		i += 1
		url := "http://api.bilibili.com/x/v2/dm/web/seg.so"
		resp, err := d.Req.R().SetQueryParams(map[string]string{"type": "1", "oid": cid, "segment_index": fmt.Sprintf("%v", i)}).Get(url)
		if err != nil {
			return result, err
		}
		if resp.StatusCode() != 200 {
			break
		} else {
			err := proto.Unmarshal(resp.Body(), &tmp)
			if err != nil {
				fmt.Println(err)
			}
			result = append(result, tmp.Elems...)
		}
	}
	return result, nil
}
