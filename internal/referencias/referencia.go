package referencias

import (
	"os"
	"path/filepath"
)

func GuardarRama(
	rutaRepositorio string,
	nombreRama string,
	hash string,
) error {

	ruta := filepath.Join(
		rutaRepositorio,
		".goforge",
		"refs",
		"heads",
		nombreRama,
	)

	return os.WriteFile(
		ruta,
		[]byte(hash),
		0644,
	)

}
