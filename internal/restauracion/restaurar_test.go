package restauracion

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/AlexanderVallejos10/goforge/internal/commits"
	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func crearCommitPrueba(
	t *testing.T,
	rutaRepositorio string,
	archivos map[string]string,
) string {

	t.Helper()

	entradas := make(
		[]indice.Entrada,
		0,
		len(archivos),
	)

	for nombre, contenido := range archivos {

		hashBlob, err := objetos.GuardarObjeto(
			rutaRepositorio,
			objetos.TipoBlob,
			[]byte(contenido),
		)

		if err != nil {
			t.Fatalf(
				"no se pudo guardar blob: %v",
				err,
			)
		}

		entradas = append(
			entradas,
			indice.Entrada{
				Archivo: nombre,
				Hash:    hashBlob,
			},
		)
	}

	contenidoTree, err := commits.CrearTree(
		entradas,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo crear tree: %v",
			err,
		)
	}

	hashTree, err := objetos.GuardarObjeto(
		rutaRepositorio,
		objetos.TipoTree,
		contenidoTree,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo guardar tree: %v",
			err,
		)
	}

	contenidoCommit, err := commits.CrearCommit(
		hashTree,
		"",
		"commit de prueba",
		"Tester",
	)

	if err != nil {
		t.Fatalf(
			"no se pudo crear commit: %v",
			err,
		)
	}

	hashCommit, err := objetos.GuardarObjeto(
		rutaRepositorio,
		objetos.TipoCommit,
		contenidoCommit,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo guardar commit: %v",
			err,
		)
	}

	return hashCommit
}

func TestRestaurarCommitCreaArchivos(
	t *testing.T,
) {

	rutaRepositorio := t.TempDir()

	hashCommit := crearCommitPrueba(
		t,
		rutaRepositorio,
		map[string]string{
			"archivo.txt": "contenido restaurado",
		},
	)

	err := RestaurarCommit(
		rutaRepositorio,
		hashCommit,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo restaurar el commit: %v",
			err,
		)
	}

	contenido, err := os.ReadFile(
		filepath.Join(
			rutaRepositorio,
			"archivo.txt",
		),
	)

	if err != nil {
		t.Fatalf(
			"no se pudo leer el archivo restaurado: %v",
			err,
		)
	}

	if string(contenido) != "contenido restaurado" {
		t.Fatalf(
			"contenido incorrecto: %q",
			string(contenido),
		)
	}
}

func TestRestaurarCommitCreaDirectorios(
	t *testing.T,
) {

	rutaRepositorio := t.TempDir()

	hashCommit := crearCommitPrueba(
		t,
		rutaRepositorio,
		map[string]string{
			filepath.Join(
				"carpeta",
				"archivo.txt",
			): "contenido interno",
		},
	)

	err := RestaurarCommit(
		rutaRepositorio,
		hashCommit,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo restaurar el commit: %v",
			err,
		)
	}

	rutaArchivo := filepath.Join(
		rutaRepositorio,
		"carpeta",
		"archivo.txt",
	)

	contenido, err := os.ReadFile(
		rutaArchivo,
	)

	if err != nil {
		t.Fatalf(
			"no se creó el archivo dentro del directorio: %v",
			err,
		)
	}

	if string(contenido) != "contenido interno" {
		t.Fatalf(
			"contenido incorrecto: %q",
			string(contenido),
		)
	}
}

func TestRestaurarCommitEliminaArchivoObsoleto(
	t *testing.T,
) {

	rutaRepositorio := t.TempDir()

	hashBlobObsoleto, err := objetos.GuardarObjeto(
		rutaRepositorio,
		objetos.TipoBlob,
		[]byte("archivo anterior"),
	)

	if err != nil {
		t.Fatalf(
			"no se pudo guardar blob anterior: %v",
			err,
		)
	}

	nombreObsoleto := "obsoleto.txt"

	err = os.WriteFile(
		filepath.Join(
			rutaRepositorio,
			nombreObsoleto,
		),
		[]byte("archivo anterior"),
		0644,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo crear archivo anterior: %v",
			err,
		)
	}

	err = indice.Guardar(
		rutaRepositorio,
		[]indice.Entrada{
			{
				Archivo: nombreObsoleto,
				Hash:    hashBlobObsoleto,
			},
		},
	)

	if err != nil {
		t.Fatalf(
			"no se pudo guardar índice anterior: %v",
			err,
		)
	}

	hashCommit := crearCommitPrueba(
		t,
		rutaRepositorio,
		map[string]string{
			"actual.txt": "archivo actual",
		},
	)

	err = RestaurarCommit(
		rutaRepositorio,
		hashCommit,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo restaurar el commit: %v",
			err,
		)
	}

	if _, err := os.Stat(
		filepath.Join(
			rutaRepositorio,
			nombreObsoleto,
		),
	); !os.IsNotExist(err) {

		t.Fatal(
			"el archivo obsoleto debería haberse eliminado",
		)
	}
}

func TestRestaurarCommitActualizaIndice(
	t *testing.T,
) {

	rutaRepositorio := t.TempDir()

	hashCommit := crearCommitPrueba(
		t,
		rutaRepositorio,
		map[string]string{
			"uno.txt": "uno",
			"dos.txt": "dos",
		},
	)

	err := RestaurarCommit(
		rutaRepositorio,
		hashCommit,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo restaurar el commit: %v",
			err,
		)
	}

	entradas, err := indice.Leer(
		rutaRepositorio,
	)

	if err != nil {
		t.Fatalf(
			"no se pudo leer el índice: %v",
			err,
		)
	}

	if len(entradas) != 2 {
		t.Fatalf(
			"se esperaban 2 entradas, se obtuvieron %d",
			len(entradas),
		)
	}

	nombres := make(
		map[string]bool,
		len(entradas),
	)

	for _, entrada := range entradas {
		nombres[entrada.Archivo] = true
	}

	if !nombres["uno.txt"] {
		t.Error(
			"falta uno.txt en el índice",
		)
	}

	if !nombres["dos.txt"] {
		t.Error(
			"falta dos.txt en el índice",
		)
	}
}
