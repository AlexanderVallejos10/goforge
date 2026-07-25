package cabeza

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

const prefijoReferencia = "ref: refs/heads/"

func Guardar(
	rutaRepositorio string,
	nombreRama string,
) error {

	if nombreRama == "" {
		return errors.New("el nombre de la rama no puede estar vacío")
	}

	ruta := filepath.Join(
		rutaRepositorio,
		".goforge",
		"HEAD",
	)

	contenido := prefijoReferencia + nombreRama

	return os.WriteFile(
		ruta,
		[]byte(contenido),
		0644,
	)
}

func Leer(
	rutaRepositorio string,
) (string, error) {

	ruta := filepath.Join(
		rutaRepositorio,
		".goforge",
		"HEAD",
	)

	datos, err := os.ReadFile(ruta)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(datos)), nil
}

func LeerRamaActual(
	rutaRepositorio string,
) (string, error) {

	contenido, err := Leer(rutaRepositorio)

	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(
		contenido,
		prefijoReferencia,
	) {
		return "", errors.New(
			"HEAD no contiene una referencia válida",
		)
	}

	nombreRama := strings.TrimPrefix(
		contenido,
		prefijoReferencia,
	)

	if nombreRama == "" {
		return "", errors.New(
			"HEAD no indica ninguna rama",
		)
	}

	return nombreRama, nil
}
