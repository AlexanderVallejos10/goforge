package indice

import (
	"path/filepath"
	"sort"
)

func Actualizar(
	anteriores []Entrada,
	nuevas []Entrada,
) []Entrada {

	porArchivo := make(
		map[string]Entrada,
		len(anteriores)+len(nuevas),
	)

	for _, entrada := range anteriores {

		ruta := filepath.Clean(
			entrada.Archivo,
		)

		entrada.Archivo = ruta
		porArchivo[ruta] = entrada
	}

	for _, entrada := range nuevas {

		ruta := filepath.Clean(
			entrada.Archivo,
		)

		entrada.Archivo = ruta
		porArchivo[ruta] = entrada
	}

	resultado := make(
		[]Entrada,
		0,
		len(porArchivo),
	)

	for _, entrada := range porArchivo {
		resultado = append(
			resultado,
			entrada,
		)
	}

	sort.Slice(
		resultado,
		func(i int, j int) bool {
			return resultado[i].Archivo < resultado[j].Archivo
		},
	)

	return resultado
}
