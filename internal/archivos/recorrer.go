package archivos

import (
	"os"
	"path/filepath"
)

func BuscarArchivos(
	ruta string,
) ([]string, error) {

	var archivos []string

	err := filepath.Walk(
		ruta,
		func(
			rutaActual string,
			info os.FileInfo,
			err error,
		) error {

			if err != nil {
				return err
			}

			if DebeIgnorar(rutaActual) {

				if info.IsDir() {
					return filepath.SkipDir
				}

				return nil
			}

			if info.IsDir() {
				return nil
			}

			if filepath.Base(rutaActual) == ".git" {
				return nil
			}

			archivos = append(
				archivos,
				rutaActual,
			)

			return nil
		},
	)

	return archivos, err

}
