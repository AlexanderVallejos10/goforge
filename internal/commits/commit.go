package commits

import (
	"encoding/json"
	"time"

	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func CrearCommit(
	hashTree string,
	hashPadre string,
	mensaje string,
	autor string,
) ([]byte, error) {

	commit := objetos.Commit{

		Tree: hashTree,

		Padre: hashPadre,

		Autor: autor,

		Mensaje: mensaje,

		Fecha: time.Now(),
	}

	return json.Marshal(
		commit,
	)

}
