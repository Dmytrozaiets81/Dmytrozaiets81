package common

import (
	"fmt"
	"github/wnstar/bili-utils/api/webInterface"
	"hash/crc32"
	"strconv"
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

func DmMidHash2Uid(midHash string) int64 {
	hashid := ""
	i := int64(0)
	for ; hashid != midHash; i++ {
		s := crc32.ChecksumIEEE([]byte(strconv.FormatInt(i, 10)))
		hashid = strconv.FormatInt(int64(s), 16)
	}
	return i
}
