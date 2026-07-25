package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/limpieza"
)

func EjecutarClean() {

	err := limpieza.Limpiar(
		".",
	)

	if err != nil {

		fmt.Println(
			"Error limpiando archivos:",
			err,
		)

		return
	}

	fmt.Println(
		"Archivos no rastreados eliminados.",
	)
}
