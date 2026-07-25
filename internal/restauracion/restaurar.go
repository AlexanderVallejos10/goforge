package restauracion

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/AlexanderVallejos10/goforge/internal/commits"
	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func RestaurarCommit(
	rutaRepositorio string,
	hashCommit string,
) error {

	if hashCommit == "" {
		return errors.New("la rama no tiene commits")
	}

	commit, err := commits.LeerCommit(
		rutaRepositorio,
		hashCommit,
	)

	if err != nil {
		return err
	}

	tree, err := commits.LeerTree(
		rutaRepositorio,
		commit.Tree,
	)

	if err != nil {
		return err
	}

	entradasIndice := make(
		[]indice.Entrada,
		0,
		len(tree.Entradas),
	)

	for _, entrada := range tree.Entradas {

		objeto, err := objetos.Leer(
			rutaRepositorio,
			entrada.Hash,
		)

		if err != nil {
			return err
		}

		if objeto.Tipo != objetos.TipoBlob {
			return errors.New(
				"el tree contiene un objeto que no es blob",
			)
		}

		rutaArchivo := filepath.Join(
			rutaRepositorio,
			entrada.Nombre,
		)

		err = os.MkdirAll(
			filepath.Dir(rutaArchivo),
			0755,
		)

		if err != nil {
			return err
		}

		err = os.WriteFile(
			rutaArchivo,
			objeto.Contenido,
			0644,
		)

		if err != nil {
			return err
		}

		entradasIndice = append(
			entradasIndice,
			indice.Entrada{
				Archivo: entrada.Nombre,
				Hash:    entrada.Hash,
			},
		)
	}

	return indice.Guardar(
		rutaRepositorio,
		entradasIndice,
	)
}
