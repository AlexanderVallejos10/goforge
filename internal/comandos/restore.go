package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/restauracion"
)

func EjecutarRestore(
	archivo string,
) {

	err := restauracion.RestaurarArchivo(
		".",
		archivo,
	)

	if err != nil {

		fmt.Println(
			"Error restaurando archivo:",
			err,
		)

		return
	}

	fmt.Println(
		"Archivo restaurado:",
		archivo,
	)
}
