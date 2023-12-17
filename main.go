package main

import (
	"fmt"

	"github.com/gabrieljacobo/godesde0/ejercicios"
)

func main() {
	numero, mensaje := ejercicios.ConvertiraInt("10")

	fmt.Println(numero)
	fmt.Println(mensaje)
}
