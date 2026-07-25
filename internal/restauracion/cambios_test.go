package restauracion

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func prepararRepositorioPrueba(
	t *testing.T,
	contenido []byte,
) (string, string) {

	t.Helper()

	rutaRepositorio := t.TempDir()
	nombreArchivo := "prueba.txt"

	err := os.MkdirAll(
		filepath.Join(
			rutaRepositorio,
			".goforge",
			"objects",
		),
		0755,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo crear el directorio de objetos: %v",
			err,
		)
	}

	hash, err := objetos.GuardarObjeto(
		rutaRepositorio,
		objetos.TipoBlob,
		contenido,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo guardar el blob: %v",
			err,
		)
	}

	err = os.WriteFile(
		filepath.Join(
			rutaRepositorio,
			nombreArchivo,
		),
		contenido,
		0644,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo crear el archivo de prueba: %v",
			err,
		)
	}

	err = indice.Guardar(
		rutaRepositorio,
		[]indice.Entrada{
			{
				Archivo: nombreArchivo,
				Hash:    hash,
			},
		},
	)

	if err != nil {
		t.Fatalf(
			"no se pudo guardar el índice: %v",
			err,
		)
	}

	return rutaRepositorio, nombreArchivo
}

func TestTieneCambiosLocalesArchivoSinCambios(
	t *testing.T,
) {

	contenido := []byte(
		"contenido original",
	)

	rutaRepositorio, _ := prepararRepositorioPrueba(
		t,
		contenido,
	)

	hayCambios, err := TieneCambiosLocales(
		rutaRepositorio,
	)

	if err != nil {
		t.Fatalf(
			"se produjo un error inesperado: %v",
			err,
		)
	}

	if hayCambios {
		t.Fatal(
			"no deberían detectarse cambios locales",
		)
	}
}

func TestTieneCambiosLocalesArchivoModificado(
	t *testing.T,
) {

	contenido := []byte(
		"contenido original",
	)

	rutaRepositorio, nombreArchivo :=
		prepararRepositorioPrueba(
			t,
			contenido,
		)

	err := os.WriteFile(
		filepath.Join(
			rutaRepositorio,
			nombreArchivo,
		),
		[]byte("contenido modificado"),
		0644,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo modificar el archivo: %v",
			err,
		)
	}

	hayCambios, err := TieneCambiosLocales(
		rutaRepositorio,
	)

	if err != nil {
		t.Fatalf(
			"se produjo un error inesperado: %v",
			err,
		)
	}

	if !hayCambios {
		t.Fatal(
			"debería detectarse el archivo modificado",
		)
	}
}

func TestTieneCambiosLocalesArchivoEliminado(
	t *testing.T,
) {

	contenido := []byte(
		"contenido original",
	)

	rutaRepositorio, nombreArchivo :=
		prepararRepositorioPrueba(
			t,
			contenido,
		)

	err := os.Remove(
		filepath.Join(
			rutaRepositorio,
			nombreArchivo,
		),
	)

	if err != nil {
		t.Fatalf(
			"no se pudo eliminar el archivo: %v",
			err,
		)
	}

	hayCambios, err := TieneCambiosLocales(
		rutaRepositorio,
	)

	if err != nil {
		t.Fatalf(
			"se produjo un error inesperado: %v",
			err,
		)
	}

	if !hayCambios {
		t.Fatal(
			"debería detectarse el archivo eliminado",
		)
	}
}

func TestTieneCambiosLocalesIgnoraArchivoNuevo(
	t *testing.T,
) {

	contenido := []byte(
		"contenido original",
	)

	rutaRepositorio, _ := prepararRepositorioPrueba(
		t,
		contenido,
	)

	err := os.WriteFile(
		filepath.Join(
			rutaRepositorio,
			"archivo-nuevo.txt",
		),
		[]byte("archivo no rastreado"),
		0644,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo crear el archivo nuevo: %v",
			err,
		)
	}

	hayCambios, err := TieneCambiosLocales(
		rutaRepositorio,
	)

	if err != nil {
		t.Fatalf(
			"se produjo un error inesperado: %v",
			err,
		)
	}

	if hayCambios {
		t.Fatal(
			"un archivo no rastreado no debería bloquear checkout",
		)
	}
}
