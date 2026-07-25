package estado

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/AlexanderVallejos10/goforge/internal/archivos"
	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func Calcular(
	rutaRepositorio string,
) ([]Cambio, error) {

	entradas, err := indice.Leer(
		rutaRepositorio,
	)

	if err != nil {
		return nil, err
	}

	indicePorArchivo := make(
		map[string]string,
		len(entradas),
	)

	for _, entrada := range entradas {

		rutaLimpia := filepath.Clean(
			entrada.Archivo,
		)

		indicePorArchivo[rutaLimpia] = entrada.Hash
	}

	archivosActuales, err := archivos.BuscarArchivos(
		rutaRepositorio,
		rutaRepositorio,
	)

	if err != nil {
		return nil, err
	}

	encontrados := make(
		map[string]bool,
		len(archivosActuales),
	)

	var cambios []Cambio

	for _, rutaArchivo := range archivosActuales {

		rutaRelativa, err := filepath.Rel(
			rutaRepositorio,
			rutaArchivo,
		)

		if err != nil {
			return nil, err
		}

		rutaRelativa = filepath.Clean(
			rutaRelativa,
		)

		encontrados[rutaRelativa] = true

		contenido, err := os.ReadFile(
			rutaArchivo,
		)

		if err != nil {
			return nil, err
		}

		contenidoObjeto := objetos.CrearContenidoObjeto(
			objetos.TipoBlob,
			contenido,
		)

		hashActual := objetos.CalcularHash(
			contenidoObjeto,
		)

		hashIndice, existe := indicePorArchivo[rutaRelativa]

		if !existe {

			cambios = append(
				cambios,
				Cambio{
					Archivo: rutaRelativa,
					Tipo:    TipoNuevo,
				},
			)

			continue
		}

		if hashActual != hashIndice {

			cambios = append(
				cambios,
				Cambio{
					Archivo: rutaRelativa,
					Tipo:    TipoModificado,
				},
			)
		}
	}

	for archivo := range indicePorArchivo {

		if encontrados[archivo] {
			continue
		}

		cambios = append(
			cambios,
			Cambio{
				Archivo: archivo,
				Tipo:    TipoEliminado,
			},
		)
	}

	sort.Slice(
		cambios,
		func(i int, j int) bool {

			return cambios[i].Archivo < cambios[j].Archivo
		},
	)

	return cambios, nil
}
