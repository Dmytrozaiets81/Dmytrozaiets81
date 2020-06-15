package common

import (
	"fmt"
	"github/wnstar/bili-utils/api/webInterface"
)

func Bv2av(bv string) string {
	resp, err := webInterface.NewViewFromId(bv)
	if err != nil {
		return ""
	}
	if resp.Code == 0 {
		return fmt.Sprintf("av%d", resp.Data.Aid)
	} else {
		return ""
	}
}

func Av2bv(av string) string {
	resp, err := webInterface.NewViewFromId(av)
	if err != nil {
		return ""
	}
	if resp.Code == 0 {
		return resp.Data.Bvid
	} else {
		return ""
	}
}
