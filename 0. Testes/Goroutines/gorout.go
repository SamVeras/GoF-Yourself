package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func BaixarArquivo(x int) {
	// Simular o tempo de download
	time.Sleep(time.Duration(rand.IntN(2000)) * time.Millisecond)
	fmt.Printf("Arquivo %v baixado.\n", x)
}

func main() {
	n := 15

	t := time.Now()
	fmt.Printf("\"Download\" de %d arquivos sem goroutines:\n", n)
	for i := range n {
		BaixarArquivo(i + 1)
	}
	fmt.Println("Tempo:", time.Since(t))

	fmt.Println("\nUtilizando threads:")
	s := time.Now()
	var wg sync.WaitGroup
	for i := range n {
		wg.Add(1)
		go func(m int) {
			defer wg.Done()
			BaixarArquivo(m + 1)
		}(i)
	}

	wg.Wait()
	fmt.Println("Tempo:", time.Since(s))
}
