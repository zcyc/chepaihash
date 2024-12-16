package main

import (
	"fmt"
	"os"

	"github.com/zcyc/chepaihash/chepaihash"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法：chepaihash <输入字符串>")
		os.Exit(1)
	}

	arg := os.Args[1]
	chepai, _ := chepaihash.Hash(arg)
	fmt.Println(chepai)
}
