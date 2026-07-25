package comandos

import (
	"fmt"
	"os"

	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func EjecutarHash(
	rutaArchivo string,
) {

	contenido, err := os.ReadFile(
		rutaArchivo,
	)

	if err != nil {

		fmt.Println(
			"No se pudo leer el archivo:",
			err,
		)

		return
	}

	hash := objetos.CalcularHash(
		contenido,
	)

	fmt.Println(
		"Hash:",
		hash,
	)

}
