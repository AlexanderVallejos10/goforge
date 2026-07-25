package restauracion

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func RestaurarArchivo(
	rutaRepositorio string,
	nombreArchivo string,
) error {

	entradas, err := indice.Leer(
		rutaRepositorio,
	)

	if err != nil {
		return err
	}

	var entradaEncontrada *indice.Entrada

	for i := range entradas {

		if filepath.Clean(
			entradas[i].Archivo,
		) == filepath.Clean(nombreArchivo) {

			entradaEncontrada = &entradas[i]
			break
		}
	}

	if entradaEncontrada == nil {

		return errors.New(
			"el archivo no está registrado en el índice",
		)
	}

	objeto, err := objetos.Leer(
		rutaRepositorio,
		entradaEncontrada.Hash,
	)

	if err != nil {
		return err
	}

	if objeto.Tipo != objetos.TipoBlob {

		return errors.New(
			"el índice no apunta a un blob",
		)
	}

	rutaArchivo := filepath.Join(
		rutaRepositorio,
		nombreArchivo,
	)

	return os.WriteFile(
		rutaArchivo,
		objeto.Contenido,
		0644,
	)
}
