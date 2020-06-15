package example

import (
	"fmt"
	"github/wnstar/bili-utils/api/dm"
	"github/wnstar/bili-utils/common"
)

func main() {
	elements := dm.GetDm("200680805")
	fmt.Println(elements)
	fmt.Println(common.Bv2av("BV1XC4y1h78J"))
	fmt.Println(common.Av2bv("795898202"))
}
