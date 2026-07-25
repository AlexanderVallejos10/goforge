package comandos

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/cabeza"
	"github.com/AlexanderVallejos10/goforge/internal/commits"
	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
	"github.com/AlexanderVallejos10/goforge/internal/referencias"
)

func EjecutarCommit(
	mensaje string,
) {

	if mensaje == "" {

		fmt.Println(
			"El mensaje del commit no puede estar vacío.",
		)

		return
	}

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

	entradas, err := indice.Leer(
		".",
	)

	if err != nil {

		fmt.Println(
			"Error leyendo índice:",
			err,
		)

		return
	}

	if len(entradas) == 0 {

		fmt.Println(
			"No hay archivos preparados para crear el commit.",
		)

		return
	}

	datosTree, err := commits.CrearTree(
		entradas,
	)

	if err != nil {

		fmt.Println(
			"Error creando tree:",
			err,
		)

		return
	}

	hashTree, err := objetos.GuardarObjeto(
		".",
		objetos.TipoTree,
		datosTree,
	)

	if err != nil {

		fmt.Println(
			"Error guardando tree:",
			err,
		)

		return
	}

	hashPadre, err := referencias.LeerRama(
		".",
		nombreRama,
	)

	if err != nil {

		fmt.Println(
			"Error leyendo la rama:",
			err,
		)

		return
	}

	datosCommit, err := commits.CrearCommit(
		hashTree,
		hashPadre,
		mensaje,
		"AlexanderVallejos10",
	)

	if err != nil {

		fmt.Println(
			"Error creando commit:",
			err,
		)

		return
	}

	hashCommit, err := objetos.GuardarObjeto(
		".",
		objetos.TipoCommit,
		datosCommit,
	)

	if err != nil {

		fmt.Println(
			"Error guardando commit:",
			err,
		)

		return
	}

	err = referencias.GuardarRama(
		".",
		nombreRama,
		hashCommit,
	)

	if err != nil {

		fmt.Println(
			"Error actualizando rama:",
			err,
		)

		return
	}

	fmt.Println(
		"Commit creado:",
		hashCommit,
	)

	fmt.Println(
		"Rama:",
		nombreRama,
	)
}
