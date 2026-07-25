package archivos

import (
	"path/filepath"
	"strings"
)

func DebeIgnorar(
	rutaRepositorio string,
	ruta string,
) bool {

	nombre := filepath.Base(
		ruta,
	)

	// Ignorados internos obligatorios

	if nombre == ".git" {
		return true
	}

	if nombre == ".goforge" {
		return true
	}

	reglas := LeerIgnorados(
		rutaRepositorio,
	)

	rutaLimpia := filepath.Clean(
		ruta,
	)

	for _, regla := range reglas {

		regla = strings.TrimSpace(
			regla,
		)

		if regla == "" {
			continue
		}

		// carpetas completas

		if strings.HasSuffix(
			regla,
			"/",
		) {

			regla = strings.TrimSuffix(
				regla,
				"/",
			)

			if nombre == regla {
				return true
			}
		}

		// coincidencia directa

		if nombre == regla {
			return true
		}

		// patrones simples (*.tmp)

		coincide, err := filepath.Match(
			regla,
			nombre,
		)

		if err == nil && coincide {
			return true
		}

		// ruta completa

		if rutaLimpia == filepath.Clean(regla) {
			return true
		}
	}

	return false
}
