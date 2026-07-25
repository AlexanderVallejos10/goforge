package objetos

import (
	"os"
	"path/filepath"
)

func LeerObjetoCompleto(
	rutaRepositorio string,
	hash string,
) ([]byte, error) {

	ruta := filepath.Join(
		rutaRepositorio,
		".goforge",
		"objects",
		hash[:2],
		hash[2:],
	)

	return os.ReadFile(
		ruta,
	)

}
