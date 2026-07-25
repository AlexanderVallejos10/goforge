package objetos

import (
	"os"
	"path/filepath"
)

func LeerObjeto(
	rutaRepositorio string,
	hash string,
) ([]byte, error) {

	rutaObjeto := filepath.Join(
		rutaRepositorio,
		".goforge",
		"objects",
		hash[:2],
		hash[2:],
	)

	return os.ReadFile(
		rutaObjeto,
	)

}
