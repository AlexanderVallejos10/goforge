package indice

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Leer(
	rutaRepositorio string,
) ([]Entrada, error) {

	rutaIndice := filepath.Join(
		rutaRepositorio,
		".goforge",
		"index",
	)

	datos, err := os.ReadFile(
		rutaIndice,
	)

	if err != nil {
		return nil, err
	}

	var entradas []Entrada

	err = json.Unmarshal(
		datos,
		&entradas,
	)

	return entradas, err

}
