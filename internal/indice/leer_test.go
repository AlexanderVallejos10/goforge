package indice

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGuardarYLeerIndice(t *testing.T) {

	rutaTemporal := t.TempDir()

	os.MkdirAll(
		filepath.Join(
			rutaTemporal,
			".goforge",
		),
		0755,
	)

	entradasOriginales := []Entrada{

		{
			Archivo: "archivo.txt",
			Hash:    "abc123",
		},
	}

	err := Guardar(
		rutaTemporal,
		entradasOriginales,
	)

	if err != nil {
		t.Fatal(err)
	}

	entradasLeidas, err := Leer(
		rutaTemporal,
	)

	if err != nil {
		t.Fatal(err)
	}

	if len(entradasLeidas) != 1 {

		t.Fatalf(
			"se esperaban 1 entrada, llegaron %d",
			len(entradasLeidas),
		)

	}

	if entradasLeidas[0].Hash != "abc123" {

		t.Fatal(
			"el hash no coincide",
		)

	}

}
