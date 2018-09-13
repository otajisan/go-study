package main

import (
	"fmt"
	"./variants"
)

var BASE_VAL int = 100

func main() {
	fmt.Println("vim-go")
	n := variants.RunVariants()
	fmt.Println(BASE_VAL + n)
}
