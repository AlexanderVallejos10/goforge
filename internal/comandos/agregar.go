package comandos

import (
	"fmt"
	"os"

	"github.com/AlexanderVallejos10/goforge/internal/archivos"
	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func EjecutarAgregar(
	ruta string,
) {

	listaArchivos, err := archivos.BuscarArchivos(
		ruta,
	)

	if err != nil {

		fmt.Println(
			"Error buscando archivos:",
			err,
		)

		return
	}

	var entradas []indice.Entrada

	for _, archivo := range listaArchivos {

		contenido, err := os.ReadFile(
			archivo,
		)

		if err != nil {

			fmt.Println(
				"No se pudo leer:",
				archivo,
			)

			continue
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

			continue
		}

		entradas = append(
			entradas,
			indice.Entrada{
				Archivo: archivo,
				Hash:    hash,
			},
		)

	}

	err = indice.Guardar(
		".",
		entradas,
	)

	if err != nil {

		fmt.Println(
			"Error guardando índice:",
			err,
		)

		return
	}

	fmt.Println(
		"Archivos agregados:",
		len(entradas),
	)

}
