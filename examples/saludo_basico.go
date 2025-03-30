package main

import (
	"fmt"

	"github.com/edalmava/saludo"
)

func main() {
	mensaje := saludo.Hello("Edalmava")
	fmt.Println(mensaje)
}
