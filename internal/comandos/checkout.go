package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/cabeza"
	"github.com/AlexanderVallejos10/goforge/internal/ramas"
	"github.com/AlexanderVallejos10/goforge/internal/referencias"
	"github.com/AlexanderVallejos10/goforge/internal/restauracion"
)

func EjecutarCheckout(
	nombreRama string,
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

	if nombreRama == ramaActual {
		fmt.Println(
			"Ya estás en la rama:",
			nombreRama,
		)
		return
	}

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
			break
		}
	}

	if !existe {
		fmt.Println(
			"La rama no existe:",
			nombreRama,
		)
		return
	}

	hayCambios, err := restauracion.TieneCambiosLocales(
		".",
	)

	if err != nil {
		fmt.Println(
			"Error comprobando cambios locales:",
			err,
		)
		return
	}

	if hayCambios {
		fmt.Println(
			"No se puede cambiar de rama.",
		)
		fmt.Println(
			"Existen cambios locales sin confirmar.",
		)
		return
	}

	hashCommit, err := referencias.LeerRama(
		".",
		nombreRama,
	)

	if err != nil {
		fmt.Println(
			"Error leyendo la rama:",
			err,
		)
		return
	}

	err = restauracion.RestaurarCommit(
		".",
		hashCommit,
	)

	if err != nil {
		fmt.Println(
			"Error restaurando archivos:",
			err,
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
