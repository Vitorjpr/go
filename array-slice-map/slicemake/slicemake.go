package main

import (
	"fmt"
)

func main(){
	slice := make([]int, 10) //tipo, tamanho do slice
	slice[9] = 12
	fmt.Println(slice)

	slice = make([]int, 10, 20) //tipo, tamanho slice, capacidade (o compilador cria um array interno com capacidade 20)
	fmt.Println(slice, len(slice), cap(slice)) //slice, tamanho e capacidade total do array interno

	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 11) // nao causara erro de compilacao
	fmt.Println(slice, len(slice), cap(slice))
}