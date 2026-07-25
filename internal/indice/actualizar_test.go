package indice

import "testing"

func TestActualizarIndiceConservaEntradas(t *testing.T) {

	anteriores := []Entrada{
		{
			Archivo: "main.go",
			Hash:    "hash-main",
		},
		{
			Archivo: "prueba.txt",
			Hash:    "hash-anterior",
		},
	}

	nuevas := []Entrada{
		{
			Archivo: "prueba.txt",
			Hash:    "hash-nuevo",
		},
	}

	resultado := Actualizar(
		anteriores,
		nuevas,
	)

	if len(resultado) != 2 {

		t.Fatalf(
			"se esperaban 2 entradas, llegaron %d",
			len(resultado),
		)
	}

	hashPrueba := ""

	for _, entrada := range resultado {

		if entrada.Archivo == "prueba.txt" {
			hashPrueba = entrada.Hash
		}
	}

	if hashPrueba != "hash-nuevo" {

		t.Fatalf(
			"se esperaba hash-nuevo, llegó %s",
			hashPrueba,
		)
	}
}
