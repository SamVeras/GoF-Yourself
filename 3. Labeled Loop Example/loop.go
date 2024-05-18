package main

import "fmt"

func main() {
	mx := [4][3]int{
		{23, 44, 54},
		{18, 16, 32},
		{64, 28, 56},
		{52, 24, 48},
	}

	for _, v := range mx {
		fmt.Println(v)
	}

	println()

	k := 52
	found := false

Loop:
	for y, v := range mx {
		for x, u := range v {
			fmt.Printf("(%d, %d) %d\n", x+1, y+1, u)
			if u == k {
				fmt.Printf("Found %d at (%d, %d)\n", u, x+1, y+1)
				found = true
				break Loop
			}
		}
	}

	if !found {
		fmt.Println("Value not found.")
	}

}
