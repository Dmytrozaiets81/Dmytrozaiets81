package tools

import (
	"fmt"
	"hash/crc32"
	"strconv"

	"github.com/wnstar/bili-utils/core"
)

type Tools struct {
	Bilibili *core.Bilibili
}

func (t Tools) Bv2av(bv string) string {
	resp, err := t.Bilibili.WebInterface.View(bv)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("av%d", resp.Aid)
}

func (t Tools) Av2bv(av string) string {
	resp, err := t.Bilibili.WebInterface.View(av)
	if err != nil {
		return ""
	}
	return resp.Bvid
}

func (t Tools) DmMidHash2Uid(midHash string) int64 {
	hashid := ""
	i := int64(0)
	for ; hashid != midHash; i++ {
		s := crc32.ChecksumIEEE([]byte(strconv.FormatInt(i, 10)))
		hashid = strconv.FormatInt(int64(s), 16)
	}
	return i
}
