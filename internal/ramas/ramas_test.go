package ramas

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func prepararDirectorioRamas(
	t *testing.T,
) string {

	t.Helper()

	rutaRepositorio := t.TempDir()

	err := os.MkdirAll(
		filepath.Join(
			rutaRepositorio,
			".goforge",
			"refs",
			"heads",
		),
		0755,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo crear el directorio de ramas: %v",
			err,
		)
	}

	return rutaRepositorio
}

func TestCrearRama(
	t *testing.T,
) {

	rutaRepositorio := prepararDirectorioRamas(t)
	hashEsperado := "abc123"

	err := Crear(
		rutaRepositorio,
		"desarrollo",
		hashEsperado,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo crear la rama: %v",
			err,
		)
	}

	contenido, err := os.ReadFile(
		RutaRama(
			rutaRepositorio,
			"desarrollo",
		),
	)

	if err != nil {
		t.Fatalf(
			"no se pudo leer la rama creada: %v",
			err,
		)
	}

	if string(contenido) != hashEsperado {
		t.Fatalf(
			"hash incorrecto: se esperaba %q y se obtuvo %q",
			hashEsperado,
			string(contenido),
		)
	}
}

func TestListarRamas(
	t *testing.T,
) {

	rutaRepositorio := prepararDirectorioRamas(t)

	ramasEsperadas := []string{
		"main",
		"desarrollo",
		"prueba",
	}

	for _, nombre := range ramasEsperadas {

		err := Crear(
			rutaRepositorio,
			nombre,
			"hash-"+nombre,
		)

		if err != nil {
			t.Fatalf(
				"no se pudo crear la rama %s: %v",
				nombre,
				err,
			)
		}
	}

	lista, err := Listar(
		rutaRepositorio,
	)

	if err != nil {
		t.Fatalf(
			"no se pudieron listar las ramas: %v",
			err,
		)
	}

	if len(lista) != len(ramasEsperadas) {
		t.Fatalf(
			"se esperaban %d ramas y se obtuvieron %d",
			len(ramasEsperadas),
			len(lista),
		)
	}

	for _, nombre := range ramasEsperadas {

		if !slices.Contains(
			lista,
			nombre,
		) {
			t.Errorf(
				"no se encontró la rama %s",
				nombre,
			)
		}
	}
}

func TestRutaRama(
	t *testing.T,
) {

	rutaRepositorio := filepath.Join(
		"repositorio",
		"ejemplo",
	)

	rutaEsperada := filepath.Join(
		rutaRepositorio,
		".goforge",
		"refs",
		"heads",
		"main",
	)

	rutaObtenida := RutaRama(
		rutaRepositorio,
		"main",
	)

	if rutaObtenida != rutaEsperada {
		t.Fatalf(
			"ruta incorrecta: se esperaba %q y se obtuvo %q",
			rutaEsperada,
			rutaObtenida,
		)
	}
}

func TestCrearRamaExistenteDevuelveError(
	t *testing.T,
) {

	rutaRepositorio := prepararDirectorioRamas(t)

	err := Crear(
		rutaRepositorio,
		"desarrollo",
		"hash-original",
	)

	if err != nil {
		t.Fatalf(
			"no se pudo crear la rama inicial: %v",
			err,
		)
	}

	err = Crear(
		rutaRepositorio,
		"desarrollo",
		"hash-nuevo",
	)

	if err == nil {
		t.Fatal(
			"se esperaba un error al crear una rama existente",
		)
	}

	contenido, err := os.ReadFile(
		RutaRama(
			rutaRepositorio,
			"desarrollo",
		),
	)

	if err != nil {
		t.Fatalf(
			"no se pudo leer la rama original: %v",
			err,
		)
	}

	if string(contenido) != "hash-original" {
		t.Fatalf(
			"la rama existente fue modificada: %q",
			string(contenido),
		)
	}
}
