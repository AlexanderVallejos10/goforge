package api

import (
	"os"
	"path/filepath"

	"github.com/AlexanderVallejos10/goforge/internal/cabeza"
	"github.com/AlexanderVallejos10/goforge/internal/estado"
	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func GetBranch(
	ruta string,
) string {

	rama, err := cabeza.LeerRamaActual(
		ruta,
	)

	if err != nil {
		return err.Error()
	}

	return rama
}

func GetStatus(
	ruta string,
) []estado.Cambio {

	cambios, err := estado.Calcular(
		ruta,
	)

	if err != nil {
		return []estado.Cambio{}
	}

	return cambios
}

func Add(
	ruta string,
) string {

	entradas := []indice.Entrada{}

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

			if info.IsDir() {

				if info.Name() == ".goforge" {
					return filepath.SkipDir
				}

				return nil
			}

			contenido, err := os.ReadFile(
				rutaActual,
			)

			if err != nil {
				return err
			}

			hash, err := objetos.GuardarObjeto(
				ruta,
				objetos.TipoBlob,
				contenido,
			)

			if err != nil {
				return err
			}

			relativo, err := filepath.Rel(
				ruta,
				rutaActual,
			)

			if err != nil {
				return err
			}

			entradas = append(
				entradas,
				indice.Entrada{
					Archivo: filepath.Clean(
						relativo,
					),
					Hash: hash,
				},
			)

			return nil
		},
	)

	if err != nil {
		return err.Error()
	}

	err = indice.Guardar(
		ruta,
		entradas,
	)

	if err != nil {
		return err.Error()
	}

	return "Archivos agregados correctamente"

}

func Commit(
	ruta string,
	mensaje string,
) string {

	return "Commit pendiente"

}
