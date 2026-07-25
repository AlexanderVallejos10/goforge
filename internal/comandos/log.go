package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/cabeza"
	"github.com/AlexanderVallejos10/goforge/internal/commits"
	"github.com/AlexanderVallejos10/goforge/internal/referencias"
)

func EjecutarLog() {

	nombreRama, err := cabeza.LeerRamaActual(
		".",
	)

	if err != nil {

		fmt.Println(
			"Error leyendo la rama actual:",
			err,
		)

		return
	}

	hash, err := referencias.LeerRama(
		".",
		nombreRama,
	)

	if err != nil {

		fmt.Println(
			"Error leyendo historial:",
			err,
		)

		return
	}

	if hash == "" {

		fmt.Println(
			"La rama no tiene commits:",
			nombreRama,
		)

		return
	}

	fmt.Println(
		"Rama:",
		nombreRama,
	)

	fmt.Println()

	for hash != "" {

		commit, err := commits.LeerCommit(
			".",
			hash,
		)

		if err != nil {

			fmt.Println(
				"Error leyendo commit:",
				err,
			)

			return
		}

		fmt.Println(
			"commit",
			hash,
		)

		fmt.Println(
			"Autor:",
			commit.Autor,
		)

		fmt.Println(
			"Mensaje:",
			commit.Mensaje,
		)

		fmt.Println(
			"Fecha:",
			commit.Fecha.Format(
				"2006-01-02 15:04:05 -07:00",
			),
		)

		fmt.Println(
			"-------------------",
		)

		hash = commit.Padre
	}
}
