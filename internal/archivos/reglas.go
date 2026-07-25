package archivos

import (
	"os"
	"path/filepath"
	"strings"
)

func LeerIgnorados(
	rutaRepositorio string,
) []string {

	ruta := filepath.Join(
		rutaRepositorio,
		".goforgeignore",
	)

	datos, err := os.ReadFile(
		ruta,
	)

	if err != nil {
		return nil
	}

	lineas := strings.Split(
		string(datos),
		"\n",
	)

	var reglas []string

	for _, linea := range lineas {

		linea = strings.TrimSpace(
			linea,
		)

		if linea == "" {
			continue
		}

		reglas = append(
			reglas,
			linea,
		)
	}

	return reglas
}
