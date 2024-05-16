package main

import (
	"fmt"
	"time"
)

func contador(x int) {
	if x == 0 {
		fmt.Println("BOOM!")
		return
	}
	fmt.Println(x)
	time.Sleep(time.Second)
	contador(x - 1)
}

func main() {
	contador(5)
}
