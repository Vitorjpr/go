package main

import "fmt"

func main(){
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[11111111111] = "Maria"
	aprovados[22222222222] = "Paulo"
	aprovados[33333333333] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[22222222222])
	delete(aprovados, 22222222222)
	fmt.Println(aprovados)
	
	fmt.Println(aprovados[22222222222])
}