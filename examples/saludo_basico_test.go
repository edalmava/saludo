package main

import (
	"testing"

	"github.com/edalmava/saludo"
)

func TestHello(t *testing.T) {
	nombre := "Edalmava"
	esperado := "Hola, Edalmava. ¡Bienvenido!"
	resultado := saludo.Hello(nombre)
	if resultado != esperado {
		t.Errorf("Hello(%s) = %s; se esperaba %s", nombre, resultado, esperado)
	}
}
