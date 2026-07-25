package indice

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Guardar(
	rutaRepositorio string,
	entradas []Entrada,
) error {

	datos, err := json.MarshalIndent(
		entradas,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	rutaIndice := filepath.Join(
		rutaRepositorio,
		".goforge",
		"index",
	)

	return os.WriteFile(
		rutaIndice,
		datos,
		0644,
	)

}
