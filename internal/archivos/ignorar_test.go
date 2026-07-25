package archivos

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDebeIgnorarArchivoTmp(
	t *testing.T,
) {

	ruta := t.TempDir()

	err := os.WriteFile(
		filepath.Join(
			ruta,
			".goforgeignore",
		),
		[]byte("*.tmp"),
		0644,
	)

	if err != nil {
		t.Fatal(err)
	}

	// Creamos el archivo temporal dentro del repositorio

	archivo := filepath.Join(
		ruta,
		"prueba.tmp",
	)

	err = os.WriteFile(
		archivo,
		[]byte("temporal"),
		0644,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !DebeIgnorar(
		archivo,
	) {

		t.Fatal(
			"debería ignorar archivos tmp",
		)
	}
}

func TestNoIgnoraArchivoNormal(
	t *testing.T,
) {

	ruta := t.TempDir()

	err := os.WriteFile(
		filepath.Join(
			ruta,
			".goforgeignore",
		),
		[]byte("*.tmp"),
		0644,
	)

	if err != nil {
		t.Fatal(err)
	}

	archivo := filepath.Join(
		ruta,
		"archivo.txt",
	)

	err = os.WriteFile(
		archivo,
		[]byte("normal"),
		0644,
	)

	if err != nil {
		t.Fatal(err)
	}

	if DebeIgnorar(
		archivo,
	) {

		t.Fatal(
			"no debería ignorar archivos normales",
		)
	}
}
