package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	factors := flag.Args()

	if len(factors) != 4 {
		fmt.Println("因子ループをしたいウマ娘を4種類指定してください")
		return
	}

	fmt.Println("----- 第一世代 -----")
	fmt.Println(factors[0] + " 親: " + factors[1] + " + " + factors[2])
	//for _, v := range factors {
	//	fmt.Println(v)
	//}
}
