package comandos

import (
	"fmt"
	"os"

	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func EjecutarGuardar(
	rutaArchivo string,
) {

	contenido, err := os.ReadFile(
		rutaArchivo,
	)

	if err != nil {

		fmt.Println(
			"Error leyendo archivo:",
			err,
		)

		return
	}

	hash, err := objetos.GuardarObjeto(
		".",
		objetos.TipoBlob,
		contenido,
	)

	if err != nil {

		fmt.Println(
			"Error guardando objeto:",
			err,
		)

		return
	}

	fmt.Println(
		"Objeto guardado:",
		hash,
	)

}
