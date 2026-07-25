package restauracion

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"

	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func TieneCambiosLocales(
	rutaRepositorio string,
) (bool, error) {

	entradas, err := indice.Leer(
		rutaRepositorio,
	)

	if err != nil {
		return false, err
	}

	for _, entrada := range entradas {

		rutaArchivo := filepath.Join(
			rutaRepositorio,
			entrada.Archivo,
		)

		contenidoActual, err := os.ReadFile(
			rutaArchivo,
		)

		if err != nil {

			if errors.Is(err, os.ErrNotExist) {
				return true, nil
			}

			return false, err
		}

		objeto, err := objetos.Leer(
			rutaRepositorio,
			entrada.Hash,
		)

		if err != nil {
			return false, err
		}

		if objeto.Tipo != objetos.TipoBlob {
			return false, errors.New(
				"el índice contiene un objeto que no es blob",
			)
		}

		if !bytes.Equal(
			contenidoActual,
			objeto.Contenido,
		) {
			return true, nil
		}
	}

	return false, nil
}
