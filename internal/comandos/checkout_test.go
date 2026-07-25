package comandos

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/AlexanderVallejos10/goforge/internal/cabeza"
	"github.com/AlexanderVallejos10/goforge/internal/commits"
	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
	"github.com/AlexanderVallejos10/goforge/internal/ramas"
	"github.com/AlexanderVallejos10/goforge/internal/referencias"
)

func prepararCheckout(
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

	hashBlob, err := objetos.GuardarObjeto(
		ruta,
		objetos.TipoBlob,
		[]byte("contenido"),
	)

	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(
		filepath.Join(
			ruta,
			"archivo.txt",
		),
		[]byte("contenido"),
		0644,
	)

	if err != nil {
		t.Fatal(err)
	}

	entradas := []indice.Entrada{

		{
			Archivo: "archivo.txt",
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
		"commit prueba checkout",
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

	err = ramas.Crear(
		ruta,
		"main",
		hashCommit,
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

func TestCheckoutCreaRamaNueva(
	t *testing.T,
) {

	_ = prepararCheckout(t)
}

func TestRamaNoExiste(
	t *testing.T,
) {

	ruta := prepararCheckout(t)

	lista, err := ramas.Listar(
		ruta,
	)

	if err != nil {
		t.Fatal(err)
	}

	for _, rama := range lista {

		if rama == "fantasma" {

			t.Fatal(
				"la rama fantasma no debería existir",
			)
		}
	}
}

func TestCheckoutMantieneHEAD(
	t *testing.T,
) {

	ruta := prepararCheckout(t)

	rama, err := cabeza.LeerRamaActual(
		ruta,
	)

	if err != nil {
		t.Fatal(err)
	}

	if rama != "main" {

		t.Fatalf(
			"se esperaba main y se obtuvo %s",
			rama,
		)
	}
}

func TestCheckoutCrearNuevaRamaCambiaHEAD(
	t *testing.T,
) {

	ruta := prepararCheckout(t)

	original, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	defer os.Chdir(original)

	err = os.Chdir(ruta)

	if err != nil {
		t.Fatal(err)
	}

	EjecutarCheckout(
		true,
		"desarrollo",
	)

	rama, err := cabeza.LeerRamaActual(
		ruta,
	)

	if err != nil {
		t.Fatal(err)
	}

	if rama != "desarrollo" {

		t.Fatalf(
			"se esperaba desarrollo y se obtuvo %s",
			rama,
		)
	}
}

func TestCheckoutBloqueaCambiosLocales(
	t *testing.T,
) {

	ruta := prepararCheckout(t)

	err := os.WriteFile(
		filepath.Join(
			ruta,
			"archivo.txt",
		),
		[]byte("cambio local"),
		0644,
	)

	if err != nil {
		t.Fatal(err)
	}

	original, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	defer os.Chdir(original)

	err = os.Chdir(ruta)

	if err != nil {
		t.Fatal(err)
	}

	EjecutarCheckout(
		false,
		"main",
	)

	rama, err := cabeza.LeerRamaActual(
		ruta,
	)

	if err != nil {
		t.Fatal(err)
	}

	if rama != "main" {

		t.Fatalf(
			"HEAD debería mantenerse en main",
		)
	}
}
