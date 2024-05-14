package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for k := range 3 {
		fmt.Println("range", k)
	}

	//fmt.Println(k)
	// k não existe fora do escopo do loop

	for {
		fmt.Println("loop")
		break
	}

	for n := range 12 {
		if n%2 == 0 {
			continue
		}
		fmt.Print(n, " ")
	}
	fmt.Println()

}
