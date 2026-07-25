package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/cabeza"
	"github.com/AlexanderVallejos10/goforge/internal/ramas"
	"github.com/AlexanderVallejos10/goforge/internal/referencias"
	"github.com/AlexanderVallejos10/goforge/internal/restauracion"
)

func EjecutarCheckout(
	crearNueva bool,
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

	// 1. Verificar cambios locales antes de cualquier operación

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

	// 2. Crear rama nueva si usamos checkout -b

	if crearNueva {

		hashActual, err := referencias.LeerRama(
			".",
			ramaActual,
		)

		if err != nil {

			fmt.Println(
				"Error leyendo commit actual:",
				err,
			)

			return
		}

		err = ramas.Crear(
			".",
			nombreRama,
			hashActual,
		)

		if err != nil {

			fmt.Println(
				"Error creando rama:",
				err,
			)

			return
		}

	} else {

		// 3. Si no es -b verificar que la rama exista

		listaRamas, err := ramas.Listar(
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

		for _, rama := range listaRamas {

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
	}

	// 4. Leer commit de la rama destino

	hashCommit, err := referencias.LeerRama(
		".",
		nombreRama,
	)

	if err != nil {

		fmt.Println(
			"Error leyendo rama destino:",
			err,
		)

		return
	}

	// 5. Restaurar archivos del commit destino

	err = restauracion.RestaurarCommit(
		".",
		hashCommit,
	)

	if err != nil {

		fmt.Println(
			"Error restaurando archivos:",
			err,
		)

		// Si fue checkout -b y falló,
		// eliminamos la rama creada para no dejar basura

		if crearNueva {

			// por ahora no eliminamos físicamente
			// porque todavía no tenemos ramas.Delete()
		}

		return
	}

	// 6. Actualizar HEAD

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
