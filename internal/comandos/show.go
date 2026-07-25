package comandos

import (
	"encoding/json"
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func EjecutarShow(
	hash string,
) {

	objeto, err := objetos.Leer(
		".",
		hash,
	)

	if err != nil {

		fmt.Println(
			"Error leyendo objeto:",
			err,
		)

		return

	}

	fmt.Println(
		"Tipo:",
		objeto.Tipo,
	)

	fmt.Println()

	var formato json.RawMessage

	err = json.Unmarshal(
		objeto.Contenido,
		&formato,
	)

	if err == nil {

		datosBonitos, _ := json.MarshalIndent(
			formato,
			"",
			"  ",
		)

		fmt.Println(
			string(datosBonitos),
		)

		return
	}

	fmt.Println(
		string(objeto.Contenido),
	)

}
