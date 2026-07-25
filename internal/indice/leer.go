package indice

import (
	"bytes"
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

		if os.IsNotExist(err) {
			return []Entrada{}, nil
		}

		return nil, err
	}

	if len(bytes.TrimSpace(datos)) == 0 {
		return []Entrada{}, nil
	}

	var entradas []Entrada

	err = json.Unmarshal(
		datos,
		&entradas,
	)

	if err != nil {
		return nil, err
	}

	return entradas, nil
}
