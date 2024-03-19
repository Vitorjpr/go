package main

import "fmt"

func main(){
	var status map[string]string
	status = make(map[string]string)

	status["Vitor"] = "Aprovado"
	status["Jessica"] = "Aprovada"

	fmt.Println(status)
	fmt.Println(status["Vitor"])
	fmt.Println(status["Jessica"])

	for nome, status := range status{
		fmt.Printf("O(a) candidato(a) %s foi %s\n", nome, status)
	}

	delete(status, "Vitor")
  fmt.Println(status)	
}