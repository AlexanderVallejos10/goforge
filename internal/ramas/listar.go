package ramas

import (
	"os"
	"path/filepath"
)

func Listar(
	rutaRepositorio string,
) ([]string, error) {

	ruta := filepath.Join(
		rutaRepositorio,
		".goforge",
		"refs",
		"heads",
	)

	archivos, err := os.ReadDir(
		ruta,
	)

	if err != nil {
		return nil, err
	}

	var ramas []string

	for _, archivo := range archivos {

		ramas = append(
			ramas,
			archivo.Name(),
		)

	}

	return ramas, nil

}
