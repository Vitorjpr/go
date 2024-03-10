package main

import "fmt"
// não existe operador ternario

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

	// se ternario existisse: return nota >= 6 ? "Aprovado" : "Reprovado"
}

func main () {
	fmt.Println(obterResultado(6.2))
}