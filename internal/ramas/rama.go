package ramas

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
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

	nombre = strings.TrimSpace(nombre)

	if nombre == "" {
		return errors.New(
			"el nombre de la rama no puede estar vacío",
		)
	}

	ruta := RutaRama(
		rutaRepositorio,
		nombre,
	)

	_, err := os.Stat(ruta)

	if err == nil {
		return errors.New(
			"la rama ya existe",
		)
	}

	if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	return os.WriteFile(
		ruta,
		[]byte(hashActual),
		0644,
	)
}
