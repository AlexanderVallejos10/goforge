package diferencias

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

type DiferenciaArchivo struct {
	Archivo string
	Lineas  []Linea
}

func CompararArchivo(
	rutaRepositorio string,
	nombreArchivo string,
) (DiferenciaArchivo, error) {

	entradas, err := indice.Leer(
		rutaRepositorio,
	)

	if err != nil {
		return DiferenciaArchivo{}, err
	}

	var entradaEncontrada *indice.Entrada

	for i := range entradas {

		if filepath.Clean(entradas[i].Archivo) ==
			filepath.Clean(nombreArchivo) {

			entradaEncontrada = &entradas[i]
			break
		}
	}

	if entradaEncontrada == nil {
		return DiferenciaArchivo{}, errors.New(
			"el archivo no está registrado en el índice",
		)
	}

	objeto, err := objetos.Leer(
		rutaRepositorio,
		entradaEncontrada.Hash,
	)

	if err != nil {
		return DiferenciaArchivo{}, err
	}

	if objeto.Tipo != objetos.TipoBlob {
		return DiferenciaArchivo{}, errors.New(
			"la entrada del índice no apunta a un blob",
		)
	}

	rutaArchivo := filepath.Join(
		rutaRepositorio,
		nombreArchivo,
	)

	contenidoActual, err := os.ReadFile(
		rutaArchivo,
	)

	if err != nil {

		if errors.Is(err, os.ErrNotExist) {

			contenidoActual = []byte{}

		} else {

			return DiferenciaArchivo{}, err
		}
	}

	return DiferenciaArchivo{
		Archivo: nombreArchivo,
		Lineas: CompararLineas(
			objeto.Contenido,
			contenidoActual,
		),
	}, nil
}

func TieneDiferencias(
	lineas []Linea,
) bool {

	for _, linea := range lineas {

		if linea.Tipo != LineaIgual {
			return true
		}
	}

	return false
}
