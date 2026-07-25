package objetos

import (
	"strings"
	"testing"
)

func TestCrearContenidoObjeto(t *testing.T) {

	contenido := []byte(
		"Hola",
	)

	resultado := CrearContenidoObjeto(
		TipoBlob,
		contenido,
	)

	texto := string(resultado)

	if !strings.Contains(
		texto,
		"blob 4",
	) {

		t.Fatal(
			"el objeto no contiene la cabecera correcta",
		)

	}

}
