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
		return errors.New(
			"la rama no tiene commits",
		)
	}

	entradasAnteriores, err := indice.Leer(
		rutaRepositorio,
	)

	if err != nil {
		return err
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

	archivosDestino := make(
		map[string]bool,
		len(tree.Entradas),
	)

	for _, entrada := range tree.Entradas {

		archivosDestino[entrada.Nombre] = true
	}

	err = eliminarArchivosObsoletos(
		rutaRepositorio,
		entradasAnteriores,
		archivosDestino,
	)

	if err != nil {
		return err
	}

	entradasNuevas := make(
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

		entradasNuevas = append(
			entradasNuevas,
			indice.Entrada{
				Archivo: entrada.Nombre,
				Hash:    entrada.Hash,
			},
		)
	}

	return indice.Guardar(
		rutaRepositorio,
		entradasNuevas,
	)
}

func eliminarArchivosObsoletos(
	rutaRepositorio string,
	entradasAnteriores []indice.Entrada,
	archivosDestino map[string]bool,
) error {

	for _, entrada := range entradasAnteriores {

		if archivosDestino[entrada.Archivo] {
			continue
		}

		rutaArchivo := filepath.Join(
			rutaRepositorio,
			entrada.Archivo,
		)

		err := os.Remove(
			rutaArchivo,
		)

		if err != nil &&
			!errors.Is(err, os.ErrNotExist) {
			return err
		}

		eliminarDirectoriosVacios(
			rutaRepositorio,
			filepath.Dir(rutaArchivo),
		)
	}

	return nil
}

func eliminarDirectoriosVacios(
	rutaRepositorio string,
	rutaDirectorio string,
) {

	rutaRaiz, err := filepath.Abs(
		rutaRepositorio,
	)

	if err != nil {
		return
	}

	rutaActual, err := filepath.Abs(
		rutaDirectorio,
	)

	if err != nil {
		return
	}

	for rutaActual != rutaRaiz {

		err = os.Remove(
			rutaActual,
		)

		if err != nil {
			return
		}

		rutaActual = filepath.Dir(
			rutaActual,
		)
	}
}
