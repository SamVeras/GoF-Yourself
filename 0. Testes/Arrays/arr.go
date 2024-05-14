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

}
