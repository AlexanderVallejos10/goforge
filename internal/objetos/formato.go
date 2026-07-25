package objetos

import (
	"fmt"
)

func CrearContenidoObjeto(
	tipo TipoObjeto,
	contenido []byte,
) []byte {

	cabecera := fmt.Sprintf(
		"%s %d\x00",
		tipo,
		len(contenido),
	)

	return append(
		[]byte(cabecera),
		contenido...,
	)

}
