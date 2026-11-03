package main

import (
	"flag"
	"fmt"
)

func main() {
	input := flag.String("input", "", "客户反馈的描述文本")
	flag.Parse()
	fmt.Println(*input)
}
