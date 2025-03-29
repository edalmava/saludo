package saludo

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hola, %v. ¡Bienvenido!", name)
	return message
}
