package commits

import (
	"encoding/json"
	"errors"

	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func LeerTree(
	rutaRepositorio string,
	hash string,
) (objetos.Tree, error) {

	objeto, err := objetos.Leer(
		rutaRepositorio,
		hash,
	)

	if err != nil {
		return objetos.Tree{}, err
	}

	if objeto.Tipo != objetos.TipoTree {
		return objetos.Tree{}, errors.New(
			"el objeto indicado no es un tree",
		)
	}

	var tree objetos.Tree

	err = json.Unmarshal(
		objeto.Contenido,
		&tree,
	)

	if err != nil {
		return objetos.Tree{}, err
	}

	return tree, nil
}
