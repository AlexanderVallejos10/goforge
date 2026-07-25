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

	if !DebeIgnorar(
		ruta,
		filepath.Join(
			ruta,
			"prueba.tmp",
		),
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

	if DebeIgnorar(
		ruta,
		filepath.Join(
			ruta,
			"archivo.txt",
		),
	) {

		t.Fatal(
			"no debería ignorar archivos normales",
		)
	}
}
