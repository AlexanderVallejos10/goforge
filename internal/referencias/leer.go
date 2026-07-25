package referencias

import (
	"os"
	"path/filepath"
)

func LeerRama(
	rutaRepositorio string,
	nombreRama string,
) (string, error) {

	ruta := filepath.Join(
		rutaRepositorio,
		".goforge",
		"refs",
		"heads",
		nombreRama,
	)

	datos, err := os.ReadFile(
		ruta,
	)

	if err != nil {

		if os.IsNotExist(err) {
			return "", nil
		}

		return "", err
	}

	return string(datos), nil

}
