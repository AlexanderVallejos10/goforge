package diferencias

import (
	"os"
	"path/filepath"

	"github.com/AlexanderVallejos10/goforge/internal/indice"
)

func ArchivosConDiferencias(
	rutaRepositorio string,
) ([]string, error) {

	entradas, err := indice.Leer(
		rutaRepositorio,
	)

	if err != nil {
		return nil, err
	}

	var archivos []string

	for _, entrada := range entradas {

		rutaArchivo := filepath.Join(
			rutaRepositorio,
			entrada.Archivo,
		)

		_, err := os.Stat(
			rutaArchivo,
		)

		if err != nil ||
			DiferenciaArchivoExiste(
				rutaRepositorio,
				entrada.Archivo,
			) {

			archivos = append(
				archivos,
				entrada.Archivo,
			)
		}
	}

	return archivos, nil
}

func DiferenciaArchivoExiste(
	rutaRepositorio string,
	archivo string,
) bool {

	resultado, err := CompararArchivo(
		rutaRepositorio,
		archivo,
	)

	if err != nil {
		return false
	}

	return TieneDiferencias(
		resultado.Lineas,
	)
}
