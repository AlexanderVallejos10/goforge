package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/estado"
)

func EjecutarEstado() {

	cambios, err := estado.Calcular(
		".",
	)

	if err != nil {

		fmt.Println(
			"No se pudo calcular el estado:",
			err,
		)

		return
	}

	if len(cambios) == 0 {

		fmt.Println(
			"No hay cambios. El directorio está limpio.",
		)

		return
	}

	fmt.Println(
		"Cambios encontrados:",
	)

	for _, cambio := range cambios {

		fmt.Printf(
			"%-12s %s\n",
			cambio.Tipo,
			cambio.Archivo,
		)
	}
}
