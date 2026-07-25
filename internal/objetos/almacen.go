package objetos

import (
	"os"
	"path/filepath"
)

func GuardarObjeto(
	rutaRepositorio string,
	tipo TipoObjeto,
	contenido []byte,
) (string, error) {

	contenidoObjeto := CrearContenidoObjeto(
		tipo,
		contenido,
	)

	hash := CalcularHash(
		contenidoObjeto,
	)

	carpetaHash := filepath.Join(
		rutaRepositorio,
		".goforge",
		"objects",
		hash[:2],
	)

	err := os.MkdirAll(
		carpetaHash,
		0755,
	)

	if err != nil {
		return "", err
	}

	rutaObjeto := filepath.Join(
		carpetaHash,
		hash[2:],
	)

	err = os.WriteFile(
		rutaObjeto,
		contenidoObjeto,
		0644,
	)

	if err != nil {
		return "", err
	}

	return hash, nil

}
