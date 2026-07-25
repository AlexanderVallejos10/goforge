package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/diferencias"
)

func EjecutarDiff(
	nombreArchivo string,
) {

	if nombreArchivo == "" {

		archivos, err := diferencias.ArchivosConDiferencias(
			".",
		)

		if err != nil {

			fmt.Println(
				"Error buscando diferencias:",
				err,
			)

			return
		}

		if len(archivos) == 0 {

			fmt.Println(
				"No hay diferencias.",
			)

			return
		}

		fmt.Println(
			"Archivos modificados:",
		)

		for _, archivo := range archivos {

			fmt.Println(
				"M",
				archivo,
			)
		}

		return
	}

	resultado, err := diferencias.CompararArchivo(
		".",
		nombreArchivo,
	)

	if err != nil {

		fmt.Println(
			"Error mostrando diferencias:",
			err,
		)

		return
	}

	if !diferencias.TieneDiferencias(
		resultado.Lineas,
	) {

		fmt.Println(
			"No hay diferencias en:",
			resultado.Archivo,
		)

		return
	}

	fmt.Println(
		"--- índice/" + resultado.Archivo,
	)

	fmt.Println(
		"+++ trabajo/" + resultado.Archivo,
	)

	fmt.Println()

	for _, linea := range resultado.Lineas {

		prefijo := "  "

		switch linea.Tipo {

		case diferencias.LineaEliminada:
			prefijo = "- "

		case diferencias.LineaAgregada:
			prefijo = "+ "
		}

		fmt.Println(
			prefijo + linea.Contenido,
		)
	}
}
