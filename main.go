package main

import (
	"fmt"

	"github.com/gabrieljacobo/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1488)

	fmt.Println(estado)
	fmt.Println(texto)
}
