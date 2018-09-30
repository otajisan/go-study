package functions

import (
    "fmt"
)

func RunFunctions() {
    fmt.Println(plus(2, 3))
    sayHello()
    fmt.Println(div(19, 7))
    // 戻り値の破棄
    _, r := div(19, 7)
    fmt.Println(r)
    // エラー処理
    result, err := handleError(0)
    fmt.Printf("handle error. result=[%d] err=[%v]\n", result, err)
}

/* 引数と戻り値 */
func plus(x, y int) int {
    return x + y
}

/* 戻り値なし */
func sayHello() {
    fmt.Println("Hello, World!")
    return
}

/* 複数の戻り値 */
func div(a, b int) (int, int) {
    q := a / b
    r := a % b
    return q, r
}

/* 関数とエラー処理 */
func handleError(a int) (int, bool) {
    if (a == 0) {
        return a, true
    }

    return a, false
}
