package main

import (
	"fmt"
	"time"
)

func main() {
	// for parecido com while
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// for tradicional
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Print("\n")
	for i := 1 ; i <= 10 ; i++ {
		if i % 2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	//laço infinito
	for {
		fmt.Println("Para sempre")
		time.Sleep(time.Second * 2)
	}
}