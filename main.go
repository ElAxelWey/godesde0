package main

import (
	"fmt"

	variables "github.com/ElAxelWey/godesde0/Variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(88)
	fmt.Println(estado)
	fmt.Println(texto)
}
