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
	_, contenido, err := objetos.LeerContenidoObjeto(
		datos,
	)

	if err != nil {

		return objetos.Commit{}, err

	}

	var commit objetos.Commit

	json.Unmarshal(
		contenido,
		&commit,
	)

	return commit, err

}
