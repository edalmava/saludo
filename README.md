# Módulo Saludo

Un módulo simple de Go que proporciona funciones para generar saludos personalizados.

## Instalación

Para instalar este módulo, utiliza el comando `go get`:

```go
go get github.com/edalmava/saludo
```

## Uso

```go
package main

import (
	"fmt"
	"github.com/tuusuario/saludo"
)

func main() {
	// Generar un saludo para "María"
	mensaje := saludo.Hello("María")
	fmt.Println(mensaje)
}
```

Este código mostrará:

```
Hola, María. ¡Bienvenido!
```

## Documentación

### Funciones

#### `Hello(name string) string`

Devuelve un saludo personalizado para la persona nombrada.

**Parámetros:**
- `name` - El nombre de la persona a saludar

**Retorna:**
- Un string con el mensaje de saludo personalizado

## Contribuir

Las contribuciones son bienvenidas. Por favor, siente libre de enviar un Pull Request.

## Licencia

Este proyecto está licenciado bajo la [Licencia MIT](LICENSE).