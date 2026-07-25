package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func EjecutarLeer(
	hash string,
) {

	contenido, err := objetos.LeerObjeto(
		".",
		hash,
	)

	if err != nil {

		fmt.Println(
			"No se pudo leer objeto:",
			err,
		)

		return
	}

	fmt.Println(
		string(contenido),
	)

}
