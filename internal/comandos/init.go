package comandos

import (
	"fmt"
	"os"

	"github.com/AlexanderVallejos10/goforge/internal/repositorio"
)

func EjecutarInit() {

	rutaActual, err := os.Getwd()

	if err != nil {
		fmt.Println("Error obteniendo ruta")
		return
	}

	err = repositorio.Crear(
		rutaActual,
	)

	if err != nil {

		fmt.Println(
			"Error creando repositorio:",
			err,
		)

		return
	}

	fmt.Println(
		"Repositorio GoForge creado correctamente",
	)

}
