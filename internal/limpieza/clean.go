package limpieza

import (
	"os"
	"path/filepath"

	"github.com/AlexanderVallejos10/goforge/internal/indice"
)

func Limpiar(
	rutaRepositorio string,
) error {

	entradas, err := indice.Leer(
		rutaRepositorio,
	)

	if err != nil {
		return err
	}

	rastreados := make(
		map[string]bool,
		len(entradas),
	)

	for _, entrada := range entradas {

		rastreados[filepath.Clean(
			entrada.Archivo,
		)] = true
	}

	return filepath.Walk(
		rutaRepositorio,
		func(
			rutaActual string,
			info os.FileInfo,
			err error,
		) error {

			if err != nil {
				return err
			}

			if info.IsDir() {

				if info.Name() == ".goforge" ||
					info.Name() == ".git" {

					return filepath.SkipDir
				}

				return nil
			}

			relativo, err := filepath.Rel(
				rutaRepositorio,
				rutaActual,
			)

			if err != nil {
				return err
			}

			relativo = filepath.Clean(
				relativo,
			)

			if rastreados[relativo] {
				return nil
			}

			return os.Remove(
				rutaActual,
			)
		},
	)
}
