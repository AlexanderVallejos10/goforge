package diferencias

import "testing"

func TestCompararLineasSinCambios(
	t *testing.T,
) {

	resultado := CompararLineas(
		[]byte("uno\ndos"),
		[]byte("uno\ndos"),
	)

	if len(resultado) != 2 {
		t.Fatalf(
			"se esperaban 2 lÃ­neas y se obtuvieron %d",
			len(resultado),
		)
	}

	for _, linea := range resultado {

		if linea.Tipo != LineaIgual {

			t.Fatalf(
				"se esperaba una lÃ­nea igual y se obtuvo %s",
				linea.Tipo,
			)
		}
	}
}

func TestCompararLineasDetectaAgregada(
	t *testing.T,
) {

	resultado := CompararLineas(
		[]byte("uno\ndos"),
		[]byte("uno\nnueva\ndos"),
	)

	encontrada := false

	for _, linea := range resultado {

		if linea.Tipo == LineaAgregada &&
			linea.Contenido == "nueva" {

			encontrada = true
		}
	}

	if !encontrada {
		t.Fatal(
			"no se detectÃ³ la lÃ­nea agregada",
		)
	}
}

func TestCompararLineasDetectaEliminada(
	t *testing.T,
) {

	resultado := CompararLineas(
		[]byte("uno\neliminada\ndos"),
		[]byte("uno\ndos"),
	)

	encontrada := false

	for _, linea := range resultado {

		if linea.Tipo == LineaEliminada &&
			linea.Contenido == "eliminada" {

			encontrada = true
		}
	}

	if !encontrada {
		t.Fatal(
			"no se detectÃ³ la lÃ­nea eliminada",
		)
	}
}

func TestCompararLineasDetectaModificacion(
	t *testing.T,
) {

	resultado := CompararLineas(
		[]byte("versiÃ³n anterior"),
		[]byte("versiÃ³n nueva"),
	)

	if len(resultado) != 2 {
		t.Fatalf(
			"se esperaban 2 diferencias y se obtuvieron %d",
			len(resultado),
		)
	}

	if resultado[0].Tipo != LineaEliminada {
		t.Fatalf(
			"la primera lÃ­nea deberÃ­a ser eliminada y fue %s",
			resultado[0].Tipo,
		)
	}

	if resultado[1].Tipo != LineaAgregada {
		t.Fatalf(
			"la segunda lÃ­nea deberÃ­a ser agregada y fue %s",
			resultado[1].Tipo,
		)
	}
}
