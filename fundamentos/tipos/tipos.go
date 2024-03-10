package main

import (
	"fmt"
	"reflect"
	"math"
)
func main() {
	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (so positivos) uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println("O tipo de i2 é", reflect.TypeOf(i2))
	fmt.Println("O valor de i2 é", i2)

	// numero reais float32 float64
	var x = float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) //float64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)
	
	//string
	s1 := "Ola eu sou uma string" //somente aspas duplas ou crase
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho de s1 é", len(s1))

	//string com multiplas linhas
	s2 := `Ola
	Eu
	Sou
	Um
	String`
	fmt.Println("O tamanho de s2 é", len(s2))

	// char??
	char := 'a' //repare nas aspas simples
	fmt.Println("O tipo de char é", reflect.TypeOf(char))

	
}