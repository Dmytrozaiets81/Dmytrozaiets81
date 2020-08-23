package dm

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/wnstar/bili-utils/request"
)

//go:generate protoc --go_out=./ bilibili.proto

type Dm struct {
}

func (d Dm) SegSo(cid string) ([]*DanmakuElem, error) {
	result := make([]*DanmakuElem, 0)
	i := 0
	for {
		var tmp DmSegMobileReply
		i += 1
		body, resp, err := request.Get(fmt.Sprintf("/x/v2/dm/web/seg.so?type=%s&oid=%s&segment_index=%d", "1", cid, i))
		if err != nil {
			return result, err
		}
		if resp.StatusCode != 200 {
			break
		} else {
			err := proto.Unmarshal(body, &tmp)
			if err != nil {
				fmt.Println(err)
			}
			result = append(result, tmp.Elems...)
		}
	}
	return result, nil
}
