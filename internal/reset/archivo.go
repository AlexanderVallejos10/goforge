package reset

import (
	"os"
	"path/filepath"

	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func RestaurarArchivo(
	rutaRepositorio string,
	archivo string,
) error {

	entradas, err := indice.Leer(
		rutaRepositorio,
	)

	if err != nil {
		return err
	}

	var hash string

	for _, entrada := range entradas {

		if filepath.Clean(
			entrada.Archivo,
		) == filepath.Clean(
			archivo,
		) {

			hash = entrada.Hash
			break
		}
	}

	if hash == "" {

		return os.ErrNotExist
	}

	objeto, err := objetos.Leer(
		rutaRepositorio,
		hash,
	)

	if err != nil {
		return err
	}

	contenido := objeto.Contenido

	rutaArchivo := filepath.Join(
		rutaRepositorio,
		archivo,
	)

	return os.WriteFile(
		rutaArchivo,
		contenido,
		0644,
	)
}
