package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/commits"
	"github.com/AlexanderVallejos10/goforge/internal/referencias"
)

func EjecutarLog() {

	hash, err := referencias.LeerRama(
		".",
		"main",
	)

	if err != nil {

		fmt.Println(
			"Error leyendo historial:",
			err,
		)

		return

	}

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
			commit.Fecha,
		)

		fmt.Println(
			"-------------------",
		)

		hash = commit.Padre

	}

}
