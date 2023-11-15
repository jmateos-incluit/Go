package main

import (
	"fmt"
)

func main2(){

	// ARRAYS
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 32
	myArray[2] = 55

	myOtherArray := [3]int {2,32,55}

	fmt.Println("Arrays")
	fmt.Println(myArray)
	fmt.Println(myOtherArray)
	
	for i := 0; i < len(myArray); i++ {
		if myArray[i]>30 {
			fmt.Println(myArray[i])
		}
	}


	/*
	//ROMPEMOS INDICE A PROPOSITO PARA DEMOSTRAR QUE EN TIEMPO DE COMPILACIÓN NO LO DETECTÓ
	var romperArray [5] int

	for i := 0; i < 6; i++ {
		romperArray[i] = i
		fmt.Println(romperArray[i])
	}
	*/

	//SLICES
	var mySlice []string
	for i := 0; i < 6; i++ {
		mySlice = append(mySlice, "Valor " + fmt.Sprint(i))
	}

	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}

	for index, valor:=range mySlice{
		fmt.Println(index, valor)
	}


	//SLICE desde ARRAY

	myNuevoSlice := myOtherArray[:] //TRAE TODO
	fmt.Println(myNuevoSlice)
	myNuevoSlice = myOtherArray[1:2] //TRAE desde la posición 1 a la 2 (no incluye)
	fmt.Println(myNuevoSlice)
	myNuevoSlice = myOtherArray[0:3] //TRAE TODO determinando desde y hasta 
	fmt.Println(myNuevoSlice)
	myNuevoSlice = myOtherArray[2:2] //TRAE NADA
	fmt.Println(myNuevoSlice)

}