package comandos

import (
	"fmt"
	"os"
	"path/filepath"

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

	var entradasNuevas []indice.Entrada

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

		rutaRelativa, err := filepath.Rel(
			".",
			archivo,
		)

		if err != nil {

			fmt.Println(
				"No se pudo calcular la ruta:",
				archivo,
			)

			continue
		}

		entradasNuevas = append(
			entradasNuevas,
			indice.Entrada{
				Archivo: filepath.Clean(
					rutaRelativa,
				),
				Hash: hash,
			},
		)
	}

	entradasAnteriores, err := indice.Leer(
		".",
	)

	if err != nil {

		fmt.Println(
			"Error leyendo índice:",
			err,
		)

		return
	}

	entradasFinales := indice.Actualizar(
		entradasAnteriores,
		entradasNuevas,
	)

	err = indice.Guardar(
		".",
		entradasFinales,
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
		len(entradasNuevas),
	)
}
