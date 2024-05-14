package main

import (
	"fmt"
	"math"
)

const s string = "string constante"

func main() {
	fmt.Println(s)
	const n = 30000000000 // Por que isso funciona
	// const n = 1234567890 // e isso não? (cannot convert untyped float const to type int64 etc.)
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(d))
	const t = 1234
	const p = 9876
	fmt.Println(fmt.Sprint(t) + fmt.Sprint(p))
}
