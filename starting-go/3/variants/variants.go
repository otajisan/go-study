package variants

import "fmt"

func RunVariants() int {
	var x, y, z int
	x, y, z = 1, 100, 1000
	// 複数の変数をまとめて定義する場合、()で括るとわかりやすい
	var (
		id int
		name string
	)
	id = 20
	name = "otajisan"

	fmt.Println(x, y, z)
	fmt.Println(id, name)

	// 変数の定義は`var foo = bar`以外にも、`foo := bar`のようにも適できる
	n := one()

	fmt.Println(n)

	return n
}

func one() int {
	return 1
}
