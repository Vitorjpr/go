package main 

import (
	"fmt"
)

func main(){
	// homogenea (mesmo tipo) e estatica (tamanho fixo)
	var notas [3]float64
	notas[0], notas[1],notas[2] = 7.8, 5.9, 9.9

  total := 0.0
	for i := 0; i < len(notas) ; i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Media %.2f\n", media)
}