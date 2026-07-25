package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/reset"
)

func EjecutarReset(
	archivo string,
) {

	err := reset.RestaurarArchivo(
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
