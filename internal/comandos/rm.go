package comandos

import (
	"fmt"
	"os"

	"github.com/AlexanderVallejos10/goforge/internal/indice"
)

func EjecutarRm(
	archivo string,
) {

	entradas, err := indice.Leer(
		".",
	)

	if err != nil {

		fmt.Println(
			"Error leyendo índice:",
			err,
		)

		return
	}

	nuevasEntradas := indice.Eliminar(
		entradas,
		archivo,
	)

	err = indice.Guardar(
		".",
		nuevasEntradas,
	)

	if err != nil {

		fmt.Println(
			"Error actualizando índice:",
			err,
		)

		return
	}

	err = os.Remove(
		archivo,
	)

	if err != nil && !os.IsNotExist(err) {

		fmt.Println(
			"Error eliminando archivo:",
			err,
		)

		return
	}

	fmt.Println(
		"Archivo eliminado:",
		archivo,
	)
}
