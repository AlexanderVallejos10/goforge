package main

import (
	"fmt"
	"os"

	"github.com/AlexanderVallejos10/goforge/internal/comandos"
	"github.com/AlexanderVallejos10/goforge/internal/version"
)

func main() {

	if len(os.Args) < 2 {

		fmt.Printf(
			"%s %s\n",
			version.NombreProyecto,
			version.NumeroVersion,
		)

		fmt.Println(
			"Usa: goforge init",
		)

		return
	}

	comando := os.Args[1]

	switch comando {

	case "init":

		comandos.EjecutarInit()

	default:

		fmt.Println(
			"Comando desconocido:",
			comando,
		)

	}

}
