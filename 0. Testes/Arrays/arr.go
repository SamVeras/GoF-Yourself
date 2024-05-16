package main

import "fmt"

func main() {
	var t [5]int
	fmt.Println("t:", t)
	t[2] = 2
	fmt.Println("t:", t)
	fmt.Println("t[2]:", t[2])
	fmt.Println("t length: ", len(t))

	p := [3]int{9, 7, 5}
	fmt.Println("p:", p)

	k := [...]int{21, 22, 23, 24, 29, 31}
	fmt.Printf("%T\n", k)

	l := [...]int{1: 23, 9: 1, 12: 22}
	fmt.Println(l)

	tt := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(tt)
	fmt.Printf("%T\n", tt)

	var lm [2][2]int
	for i := range 2 {
		for j := range 2 {
			lm[i][j] = i + j
		}
	}

	fmt.Println(lm)

}
