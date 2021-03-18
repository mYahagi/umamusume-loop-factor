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

	fmt.Println(factors[0])
	printFather(factors[1])
	printMother(factors[2])

	fmt.Println("----- 以降ここから繰り返し -----")

	fmt.Println(factors[3])
	printFather(factors[0])
	printFatherAncestors(factors[1], factors[2])
	printMother(factors[2])

	fmt.Println(factors[2])
	printFather(factors[0])
	printFatherAncestors(factors[1], factors[2])
	printMother(factors[3])
	printMotherAncestors(factors[0], factors[2])

	fmt.Println(factors[1])
	printFather(factors[2])
	printFatherAncestors(factors[0], factors[3])
	printMother(factors[3])
	printMotherAncestors(factors[0], factors[2])

	fmt.Println(factors[0])
	printFather(factors[1])
	printFatherAncestors(factors[2], factors[3])
	printMother(factors[2])
	printMotherAncestors(factors[0], factors[3])
	fmt.Println("----- ここまで繰り返し -----")
}

func printFather(father string) {
	fmt.Println("   ├─  " + father)
}

func printFatherAncestors(grantFather, grantMother string) {
	fmt.Println("   │         ├─" + grantFather)
	fmt.Println("   │         └─" + grantMother)
}

func printMother(mother string) {
	fmt.Println("   └─  " + mother)
}

func printMotherAncestors(grantFather, grantMother string) {
	fmt.Println("             ├─" + grantFather)
	fmt.Println("             └─" + grantMother)
}
