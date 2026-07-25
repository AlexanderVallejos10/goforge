package reset

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/AlexanderVallejos10/goforge/internal/cabeza"
	"github.com/AlexanderVallejos10/goforge/internal/commits"
	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
	"github.com/AlexanderVallejos10/goforge/internal/referencias"
)

func crearRepositorioPrueba(
	t *testing.T,
) string {

	t.Helper()

	ruta := t.TempDir()

	err := os.MkdirAll(
		filepath.Join(
			ruta,
			".goforge",
			"refs",
			"heads",
		),
		0755,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = os.MkdirAll(
		filepath.Join(
			ruta,
			".goforge",
			"objects",
		),
		0755,
	)

	if err != nil {
		t.Fatal(err)
	}

	contenidoOriginal := []byte(
		"contenido original",
	)

	hashBlob, err := objetos.GuardarObjeto(
		ruta,
		objetos.TipoBlob,
		contenidoOriginal,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(
		filepath.Join(
			ruta,
			"prueba.txt",
		),
		contenidoOriginal,
		0644,
	)

	if err != nil {
		t.Fatal(err)
	}

	entradas := []indice.Entrada{
		{
			Archivo: "prueba.txt",
			Hash:    hashBlob,
		},
	}

	err = indice.Guardar(
		ruta,
		entradas,
	)

	if err != nil {
		t.Fatal(err)
	}

	datosTree, err := commits.CrearTree(
		entradas,
	)

	if err != nil {
		t.Fatal(err)
	}

	hashTree, err := objetos.GuardarObjeto(
		ruta,
		objetos.TipoTree,
		datosTree,
	)

	if err != nil {
		t.Fatal(err)
	}

	datosCommit, err := commits.CrearCommit(
		hashTree,
		"",
		"commit prueba reset hard",
		"tester",
	)

	if err != nil {
		t.Fatal(err)
	}

	hashCommit, err := objetos.GuardarObjeto(
		ruta,
		objetos.TipoCommit,
		datosCommit,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = referencias.GuardarRama(
		ruta,
		"main",
		hashCommit,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = cabeza.Guardar(
		ruta,
		"main",
	)

	if err != nil {
		t.Fatal(err)
	}

	return ruta
}

func TestRestaurarHardRestauraArchivo(
	t *testing.T,
) {

	ruta := crearRepositorioPrueba(
		t,
	)

	err := os.WriteFile(
		filepath.Join(
			ruta,
			"prueba.txt",
		),
		[]byte("contenido modificado"),
		0644,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = RestaurarHard(
		ruta,
	)

	if err != nil {
		t.Fatal(err)
	}

	contenido, err := os.ReadFile(
		filepath.Join(
			ruta,
			"prueba.txt",
		),
	)

	if err != nil {
		t.Fatal(err)
	}

	if string(contenido) != "contenido original" {

		t.Fatalf(
			"se esperaba contenido original y se obtuvo %s",
			string(contenido),
		)
	}
}
