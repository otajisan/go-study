package variants

import (
    "fmt"
)

// パッケージ変数
var n = 100

func RunPackageVariants() {
    n = n % 30
    fmt.Printf("n=%d\n", n)
    add3()
    fmt.Printf("n=%d\n", n)
}

func add3() {
    n = n + 1
}
