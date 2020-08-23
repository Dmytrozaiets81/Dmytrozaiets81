package main

import (
	"fmt"
	"github.com/wnstar/bili-utils/core"
	"github.com/wnstar/bili-utils/tools"
)

func main() {
	bilibili := core.NewBilibili()
	utils := tools.Tools{bilibili}
	elements, _ := bilibili.Dm.SegSo("200680805")
	fmt.Println(elements)
	fmt.Println(bilibili.Relation.Followers("1110510",1))
	fmt.Println(utils.Bv2av("BV1XC4y1h78J"))
	fmt.Println(utils.Av2bv("795898202"))
	fmt.Println(utils.DmMidHash2Uid("a9ce8ef5"))
}
