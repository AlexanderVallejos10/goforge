package repositorio

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Configuracion struct {
	Nombre  string `json:"nombre"`
	Version string `json:"version"`
}

func Crear(rutaProyecto string) error {

	rutaBase := filepath.Join(
		rutaProyecto,
		NombreCarpeta,
	)

	carpetas := []string{

		rutaBase,

		filepath.Join(
			rutaBase,
			CarpetaObjetos,
		),

		filepath.Join(
			rutaBase,
			CarpetaReferencias,
			CarpetaRamas,
		),

		filepath.Join(
			rutaBase,
			CarpetaReferencias,
			CarpetaEtiquetas,
		),
	}

	for _, carpeta := range carpetas {

		err := os.MkdirAll(
			carpeta,
			0755,
		)

		if err != nil {
			return err
		}
	}

	configuracion := Configuracion{
		Nombre:  "GoForge",
		Version: "1",
	}

	datos, err := json.MarshalIndent(
		configuracion,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	err = os.WriteFile(
		filepath.Join(
			rutaBase,
			ArchivoConfiguracion,
		),
		datos,
		0644,
	)

	if err != nil {
		return err
	}

	err = os.WriteFile(
		filepath.Join(
			rutaBase,
			ArchivoCabeza,
		),
		[]byte("ref: refs/heads/main"),
		0644,
	)

	if err != nil {
		return err
	}

	err = os.WriteFile(
		filepath.Join(
			rutaBase,
			ArchivoIndice,
		),
		[]byte(""),
		0644,
	)

	return err
}
