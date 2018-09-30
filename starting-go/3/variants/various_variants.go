package variants

import (
    "fmt"
    "math"
)

func RunVariousVariants() {
    // 論理値
    var b bool
    b = true
    fmt.Println(b)

    // 符号付き整数
    var i8 int8 = 127
    fmt.Println(i8)
    var i64 int64 = -100000000000000000
    fmt.Println(i64)
    // 符号なし整数
    var ui64 = 100000000000000000
    fmt.Println(ui64)
    // キャスト
    var n int32 = 12000
    i16 := int16(n)
    fmt.Println(i16)
    // 以下はうまくいかない(範囲外)
    foo := int8(n)
    fmt.Println(foo)
    // 最大値
    fmt.Printf("int8 min value = %d\n", math.MinInt8)
    fmt.Printf("int8 max value = %d\n", math.MaxInt8)
    fmt.Printf("uint32 max value = %d\n", math.MaxUint32)
    // 浮動小数点
    f32 := float32(1.0)
    fmt.Println(f32)
    zero := 0.0
    fmt.Printf("infinity + value, %f\n", 1.0 / zero)
    fmt.Printf("infinity - value, %f\n", -1.0 / zero)
    fmt.Printf("NaN value, %f\n", zero / zero)
    // 浮動小数点の精度
    // !!!原則、float32は利用しないのがオススメ
    fmt.Printf("float val = %v\n", 1.0000000000000000)
    fmt.Printf("float val = %v\n", 1.0000000000000001)
    fmt.Printf("float val = %v\n", 1.0000000000000002)
    fmt.Printf("float val = %v\n", 1.0000000000000003)
    fmt.Printf("float val = %v\n", 1.0000000000000004)
    fmt.Printf("float val = %v\n", float32(1.0000000000000004))
    fmt.Printf("float val = %v\n", float32(1.0) / float32(3.0))
    fmt.Printf("float val = %v\n", float64(1.0) / float64(3.0))
    // 複素数型
    c := 1.0 + 3i
    fmt.Printf("complex val = %v\n", c)
    cx := complex(1.0, 3)
    fmt.Printf("complex val = %v\n", cx)
    fmt.Printf("complex val (real) = %v\n", real(cx))
    fmt.Printf("complex val (imag) = %v\n", imag(cx))
    // rune型(文字(Unicodeコードポイントを表す特殊な整数型))
    r := '豪'
    fmt.Printf("rune val %v\n", r)
    // 文字列型
    str := "文字列だお"
    fmt.Printf("string val %v\n", str)
    // RAW文字列リテラル(バッククォート)
    raw := `
    ======================
    RAW文字列リテラル:
    ヒアドキュメントっぽい
    複数行の記述が
    書ける
    ======================
    `
    fmt.Printf("%v\n", raw)
    // 配列型
    arr := [3]int{1, 2, 3}
    arr[1] = 4
    fmt.Printf("array val => %v\n", arr)
    fmt.Printf("array val[2] => %v\n", arr[1])
    fmt.Printf("array val[2] => %v\n", arr[2])
    // 以下はコンパイルエラー(out of bounds)
    // fmt.Printf("array val[3] => %v\n", arr[3])
    // 要素数の省略
    var_arr := [...]int{4, 5, 6}
    // 可変長というわけではないので、以下はコンパイルエラー
    // var_arr[3] = 4
    fmt.Printf("array val => %v\n", var_arr)
    // 配列の代入(※要素数とデータ型が同一となる必要がある)
    arr = var_arr
    fmt.Printf("array val => %v\n", arr)
    // interface型
    var x interface{}
    fmt.Printf("interface val {nil expected} => %#v\n", x)
    x = 1
    fmt.Printf("interface val => %v\n", x)
    x = 3.14
    fmt.Printf("interface val => %v\n", x)
    x = '山'
    fmt.Printf("interface val => %v\n", x)
    x = "こんにちは"
    fmt.Printf("interface val => %v\n", x)
    x = 100
    // 以下はコンパイルエラー(interface型で宣言した変数を演算対象とすることはできない)
    // x = x + 200
}
