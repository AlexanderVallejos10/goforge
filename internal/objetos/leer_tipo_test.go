package objetos

import (
	"testing"
)

func TestLeerContenidoObjeto(t *testing.T) {

	datos := CrearContenidoObjeto(
		TipoCommit,
		[]byte(
			"{mensaje:test}",
		),
	)

	tipo, contenido, err := LeerContenidoObjeto(
		datos,
	)

	if err != nil {
		t.Fatal(err)
	}

	if tipo != TipoCommit {

		t.Fatalf(
			"se esperaba commit, llegó %s",
			tipo,
		)

	}

	if string(contenido) != "{mensaje:test}" {

		t.Fatal(
			"contenido incorrecto",
		)

	}

}
