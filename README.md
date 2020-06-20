# bilibili工具箱

##不断更新ing...


## 弹幕获取
```go
import "github/wnstar/bili-utils/api/dm"
cid:="200680805"
elements:=dm.GetDm(cid)
fmt.Println(elements)
```

## av转bv bv转av
```go
import "github/wnstar/bili-utils/common"
fmt.Println(common.Bv2av("BV1XC4y1h78J"))
fmt.Println(common.Av2bv("795898202"))
```

## 弹幕midHash转用户uid
```go
import	"github/wnstar/bili-utils/common"
fmt.Println(common.DmMidHash2Uid("a9ce8ef5"))
```