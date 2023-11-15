package main

import (
	"fmt"

	"rsc.io/quote"
)




func main1(){
	
	var miTexto1 string = "Hola Mundo1"
	var miTexto2 string
	var miTexto3 = "Hola Mundo3"

	var miInt1 int = 10
	var miInt2 int
	var miInt3 = 9

	var miFloat1 float64 = 6.8
	var miFloat2 float64
	var miFloat3 = 6.9

	var miBool bool = true


	fmt.Println("STRING")
	fmt.Println(miTexto1)
	fmt.Println(miTexto2)
	fmt.Println(miTexto3)
	fmt.Println("INT32")
	fmt.Println(miInt1)
	fmt.Println(miInt2)
	fmt.Println(miInt3)
	fmt.Println("FLOAT64")
	fmt.Println(miFloat1)
	fmt.Println(miFloat2)
	fmt.Println(miFloat3)
	fmt.Println("BOOLEAN")
	fmt.Println(miBool)

	fmt.Println("CONCATENACION")
	fmt.Println(miTexto1+miTexto3)
	fmt.Println(miTexto1,miInt1)
	fmt.Println(miTexto1+fmt.Sprint(miInt1))

		
	fmt.Println(quote.Hello())
}