# ChepaiHash

_ChepaiHash_ 是一个将字符串转换为中国车牌号的哈希工具

## 作为 Cli 使用

```bash
go install github.com/zcyc/chepaihash
```

```bash
chepaihash helloworld
# 湘L·3J9WF
```

## 作为 Lib 使用

1. install

```shell
go get github.com/zcyc/chepaihash
```

2. use

```go
package main

import (
	"fmt"
	"github.com/zcyc/chepaihash/chepaihash"
)

func main() {
	chepai, _ := chepaihash.Hash("helloworld")
	fmt.Println(chepai) // 湘L·3J9WF
}
```

## 为什么

我需要一个可读性强、易于识别的哈希值来匿名化数据

## 相关项目

- [cunzaizhuyi/hashplate-cn](https://github.com/cunzaizhuyi/hashplate-cn)
- [hugoattal/hashplate](https://github.com/hugoattal/hashplate)
- [jerryshell/chepaihash](https://github.com/jerryshell/chepaihash)

## 开源协议

[MIT](LICENSE)
