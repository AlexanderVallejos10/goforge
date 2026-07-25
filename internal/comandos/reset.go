package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/reset"
)

func EjecutarReset(
	argumento string,
) {

	if argumento == "--hard" {

		err := reset.RestaurarHard(
			".",
		)

		if err != nil {

			fmt.Println(
				"Error ejecutando reset --hard:",
				err,
			)

			return
		}

		fmt.Println(
			"Repositorio restaurado al último commit.",
		)

		return
	}

	err := reset.RestaurarArchivo(
		".",
		argumento,
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
		argumento,
	)
}
