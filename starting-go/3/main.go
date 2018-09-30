package main

import (
	"fmt"
	"./variants"
    "./functions"
)

var BASE_VAL int = 100

func main() {
    fmt.Println("########### 3.1 文")
	fmt.Println("vim-go")
	n := variants.RunVariants()
	fmt.Println(BASE_VAL + n)

    fmt.Println("########### 3.6 変数")
    variants.RunPackageVariants()
    variants.RunVariousVariants()

    fmt.Println("########### 3.11 関数")
    functions.RunFunctions()
}
