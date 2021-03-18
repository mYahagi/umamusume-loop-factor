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

	fmt.Println("F1: " + factors[0] + " 親: " + factors[1] + " + " + factors[2])
	fmt.Println("----- 以降ここから繰り返し -----")
	fmt.Println("F1: " + factors[3] + " 親: F1 " + factors[0] + " + " + factors[1])
	fmt.Println("F1: " + factors[2] + " 親: F1 " + factors[0] + " + " + "F1 " + factors[3])
	fmt.Println("F1: " + factors[1] + " 親: F1 " + factors[2] + " + " + "F1 " + factors[3])
	fmt.Println("F2: " + factors[0] + " 親: F1 " + factors[2] + " + " + "F1 " + factors[3])
	fmt.Println("----- ここまで繰り返し -----")
	fmt.Println("F2: " + factors[3] + " 親: F2 " + factors[0] + " + " + "F1 " + factors[1])
}
