package reset

import (
	"errors"

	"github.com/AlexanderVallejos10/goforge/internal/cabeza"
	"github.com/AlexanderVallejos10/goforge/internal/referencias"
	"github.com/AlexanderVallejos10/goforge/internal/restauracion"
)

func RestaurarHard(
	rutaRepositorio string,
) error {

	ramaActual, err := cabeza.LeerRamaActual(
		rutaRepositorio,
	)

	if err != nil {
		return err
	}

	hashCommit, err := referencias.LeerRama(
		rutaRepositorio,
		ramaActual,
	)

	if err != nil {
		return err
	}

	if hashCommit == "" {
		return errors.New(
			"la rama no tiene commits",
		)
	}

	return restauracion.RestaurarCommit(
		rutaRepositorio,
		hashCommit,
	)
}
