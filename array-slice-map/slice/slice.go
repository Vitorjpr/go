package main

import (
	"fmt"
	"reflect"
)

func main(){
	array := [3]int{1, 2, 3} //array (define tamanho fixo)
	slice := []int{1, 2, 3} //slice (tamanho variavel)

	fmt.Println(array, slice)
	fmt.Println(reflect	.TypeOf(array), reflect.TypeOf(slice))

	array2 := [5]int{1, 2, 3, 4, 5}

	//slice nao e um array, slice define um pedaço de um array

	slice2 := array2[1:3] //slice2 recebe os valores do array2 do indice 1 ao 3, mas nao incluindo o indice 3

	fmt.Println(array2, slice2)

	slice3 := array2[:2] //novo slice, que pega os valores nos indices 0 ao 2, mas nao incluindo o 2
  
	fmt.Println(array2, slice3)

	slice4 := slice2[:1]

	fmt.Println(array2, slice4)

	array2[1] = 6 // como slice eh um ponteiro, se eu mudar os valores do array base, os valores do slice tambem mudam

	fmt.Println(array2, slice2, slice3, slice4)

	
}