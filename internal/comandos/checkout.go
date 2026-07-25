package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/cabeza"
	"github.com/AlexanderVallejos10/goforge/internal/ramas"
)

func EjecutarCheckout(
	nombreRama string,
) {

	lista, err := ramas.Listar(
		".",
	)

	if err != nil {

		fmt.Println(
			"Error leyendo ramas:",
			err,
		)

		return
	}

	existe := false

	for _, rama := range lista {

		if rama == nombreRama {

			existe = true

		}

	}

	if !existe {

		fmt.Println(
			"La rama no existe:",
			nombreRama,
		)

		return

	}

	err = cabeza.Guardar(
		".",
		nombreRama,
	)

	if err != nil {

		fmt.Println(
			"Error cambiando HEAD:",
			err,
		)

		return
	}

	fmt.Println(
		"Cambiado a rama:",
		nombreRama,
	)

}
