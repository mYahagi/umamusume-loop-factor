package main

import (
	"flag"
	"fmt"
)

type parents struct {
	father string
	mother string
}

func main() {
	flag.Parse()
	factors := flag.Args()

	if len(factors) != 4 {
		fmt.Println("因子ループをしたいウマ娘を4種類指定してください")
		return
	}

	pedigree := map[string]parents{}

	fmt.Println("----- 子->親の組み合わせパターン構築開始 -----")

	pedigree[factors[0]] = parents{
		father: factors[1],
		mother: factors[2],
	}
	printPedigree(factors[0], pedigree)

	pedigree[factors[3]] = parents{
		father: factors[0],
		mother: factors[1],
	}
	printPedigree(factors[3], pedigree)

	pedigree[factors[2]] = parents{
		father: factors[0],
		mother: factors[3],
	}
	printPedigree(factors[2], pedigree)

	pedigree[factors[1]] = parents{
		father: factors[2],
		mother: factors[3],
	}
	printPedigree(factors[1], pedigree)

	printPedigree(factors[0], pedigree)

	fmt.Println("----- 組み合わせパターン構築完了 -----")
}

func printPedigree(target string, pedigree map[string]parents) {
	fmt.Println(target)

	printFather(pedigree[target].father)
	value, isThere := pedigree[pedigree[target].father]
	if isThere {
		printFatherAncestors(value.father, value.mother)
	}

	printMother(pedigree[target].mother)
	value, isThere = pedigree[pedigree[target].mother]
	if isThere {
		printMotherAncestors(value.father, value.mother)
	}
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
