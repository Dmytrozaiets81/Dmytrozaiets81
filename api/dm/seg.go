package dm

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/wnstar/gorequests"
	"github/wnstar/bili-utils/common"
)

//go:generate protoc --go_out=./ bilibili.proto

func (d DanmakuElem) Uid() int64 {
	return common.DmMidHash2Uid(d.MidHash)
}

func GetDm(cid string) []*DanmakuElem {
	result := make([]*DanmakuElem, 0)
	i := 0
	for {
		var tmp DmSegMobileReply
		i += 1
		resp, err := gorequests.
			Get("https://api.bilibili.com/x/v2/dm/web/seg.so").
			AddParams("type", "1").
			AddParams("oid", cid).
			AddParams("segment_index", fmt.Sprintf("%d", i)).
			Start()
		if err != nil {
			return result
		}
		fmt.Println(resp.Response.StatusCode)
		if resp.Response.StatusCode != 200 {
			break
		} else {
			err := proto.Unmarshal(resp.Data(), &tmp)
			if err != nil {
				fmt.Println(err)
			}
			result = append(result, tmp.Elems...)
		}
	}
	return result
}
