package ramas

import (
	"os"
	"path/filepath"
)

func RutaRama(
	rutaRepositorio string,
	nombre string,
) string {

	return filepath.Join(
		rutaRepositorio,
		".goforge",
		"refs",
		"heads",
		nombre,
	)

}

func Crear(
	rutaRepositorio string,
	nombre string,
	hashActual string,
) error {

	return os.WriteFile(
		RutaRama(
			rutaRepositorio,
			nombre,
		),
		[]byte(hashActual),
		0644,
	)

}
