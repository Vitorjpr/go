package main

import "fmt"

func main(){
	funcsPorLetra := map[string] map[string] float64{
		"G": {
			"Gabriel": 1.5,
			"Guilherme": 1.5,
		},
		"J": {
			"Joao": 1.5,
			"Jose": 1.5,
		},
		"M": {
			"Marcelo": 1.5,
			"Maria": 1.5,
		},
	}

	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "M")

	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		} 
	}
}