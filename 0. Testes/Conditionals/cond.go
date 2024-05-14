package main

import (
	"fmt"
	"time"
)

func main() {
	if 13%2 == 0 {
		fmt.Println("13 é par kk")
	} else {
		fmt.Println("13 é ímpar")
	}

	if x := 22; x < 3 {
		fmt.Println("queijo")
	}

	i := 3
	fmt.Print(i, " extenso é ")
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	fmt.Print("Hoje é ")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("fim de semana :D")
	default:
		fmt.Println("dia de semana :(")
	}

	t := time.Now()
	switch { // case sem expressão é tipo um if else diferente ? kk
	case t.Day()%2 == 0: // case não-constante daorinha
		fmt.Println("Dia par !")
	default:
		fmt.Println("Dia ímpar !")

	}

}
