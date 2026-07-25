package commits

import (
	"encoding/json"

	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func LeerCommit(
	rutaRepositorio string,
	hash string,
) (objetos.Commit, error) {

	datos, err := objetos.LeerObjetoCompleto(
		rutaRepositorio,
		hash,
	)

	if err != nil {

		return objetos.Commit{}, err
	}

	_, contenido, err := objetos.LeerContenidoObjeto(
		datos,
	)

	if err != nil {

		return objetos.Commit{}, err
	}

	var commit objetos.Commit

	err = json.Unmarshal(
		contenido,
		&commit,
	)

	if err != nil {

		return objetos.Commit{}, err
	}

	return commit, nil
}
