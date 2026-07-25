package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/cabeza"
	"github.com/AlexanderVallejos10/goforge/internal/ramas"
	"github.com/AlexanderVallejos10/goforge/internal/referencias"
)

func EjecutarBranch(
	nombre string,
) {

	ramaActual, err := cabeza.LeerRamaActual(
		".",
	)

	if err != nil {

		fmt.Println(
			"Error leyendo la rama actual:",
			err,
		)

		return
	}

	if nombre == "" {

		lista, err := ramas.Listar(
			".",
		)

		if err != nil {

			fmt.Println(
				"Error listando ramas:",
				err,
			)

			return
		}

		for _, rama := range lista {

			marca := "  "

			if rama == ramaActual {
				marca = "* "
			}

			fmt.Println(
				marca + rama,
			)
		}

		return
	}

	hashActual, err := referencias.LeerRama(
		".",
		ramaActual,
	)

	if err != nil {

		fmt.Println(
			"Error leyendo la rama actual:",
			err,
		)

		return
	}

	err = ramas.Crear(
		".",
		nombre,
		hashActual,
	)

	if err != nil {

		fmt.Println(
			"Error creando rama:",
			err,
		)

		return
	}

	fmt.Println(
		"Rama creada:",
		nombre,
	)
}
