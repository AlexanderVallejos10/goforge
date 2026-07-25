package indice

import (
	"path/filepath"
)

func Eliminar(
	entradas []Entrada,
	archivo string,
) []Entrada {

	var resultado []Entrada

	archivo = filepath.Clean(
		archivo,
	)

	for _, entrada := range entradas {

		if filepath.Clean(
			entrada.Archivo,
		) == archivo {

			continue
		}

		resultado = append(
			resultado,
			entrada,
		)
	}

	return resultado
}
