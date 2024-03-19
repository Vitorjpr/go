package main

import (
	"fmt"
)

func main(){
	funcionarios := map[string]float64{
		"Joao": 5000.0,
		"Maria": 2000.0,
		"Pedro": 3000.0, // o ultimo elemento precisa ter a virgula
	}

	funcionarios["Vitor"] = 6000.0

	delete(funcionarios, "inexistente") //se nao existe nao da problema

	for nome, salario := range funcionarios {
		fmt.Printf("Nome: %s - Salario: %.2f\n", nome, salario)
	}
}