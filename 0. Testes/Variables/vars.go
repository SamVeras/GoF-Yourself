package main

import "fmt"

func main() {
	var a = "começo"
	fmt.Println(a)

	var b, c int = 91, 7 // Isso especifica que b E c são int, ou só c?
	fmt.Println(b, "/", c, "=", b/c)

	var m, n float32 = 11.11, 12
	fmt.Printf("m(%f) is %T\n", m, m)
	fmt.Printf("n(%f) is %T\n", n, n)

	var d = true // tipo inferido
	fmt.Println(d, "is", d)

	var e int // Começa com 0 pelo visto
	fmt.Println(e)

	var f string // String inicializa vazia, blz
	fmt.Println(f)

	apl := "apple" // ta
	fmt.Println(apl)
}
