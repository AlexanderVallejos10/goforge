package objetos

import (
	"bytes"
	"errors"
)

func LeerContenidoObjeto(
	datos []byte,
) (TipoObjeto, []byte, error) {

	separador := []byte{0}

	posicion := bytes.Index(
		datos,
		separador,
	)

	if posicion == -1 {

		return "",
			nil,
			errors.New(
				"objeto sin cabecera válida",
			)

	}

	cabecera := string(
		datos[:posicion],
	)

	contenido := datos[posicion+1:]

	var tipo TipoObjeto

	for i, caracter := range cabecera {

		if caracter == ' ' {

			tipo = TipoObjeto(
				cabecera[:i],
			)

			break
		}
	}

	return tipo, contenido, nil

}
