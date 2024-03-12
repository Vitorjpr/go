package main

import (
	"fmt"
	"time"
)

func main(){
	t := time.Now()
	switch { //<- isso é um switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}

	fmt.Println(notaParaConceito(6.9))
}

func notaParaConceito(n float64) string {
	i := int(n)
	switch {
	case i >=9 && i <= 10:
		return "A"
	case i < 8:
		return "B"
	case i < 5:
		return "C"
	case i < 3:
		return "D"
	default:
		return "E"
	}
}