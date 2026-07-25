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
			"Comandos: init, hash, guardar, leer, add, status",
		)

		return
	}

	comando := os.Args[1]

	switch comando {

	case "init":

		comandos.EjecutarInit()

	case "hash":

		if len(os.Args) < 3 {

			fmt.Println(
				"Usa: goforge hash archivo",
			)

			return
		}

		comandos.EjecutarHash(
			os.Args[2],
		)

	case "guardar":

		if len(os.Args) < 3 {

			fmt.Println(
				"Usa: goforge guardar archivo",
			)

			return
		}

		comandos.EjecutarGuardar(
			os.Args[2],
		)

	case "leer":

		if len(os.Args) < 3 {

			fmt.Println(
				"Usa: goforge leer hash",
			)

			return
		}

		comandos.EjecutarLeer(
			os.Args[2],
		)

	case "add":

		if len(os.Args) < 3 {

			fmt.Println(
				"Usa: goforge add archivo-o-carpeta",
			)

			return
		}

		comandos.EjecutarAgregar(
			os.Args[2],
		)

	case "status":

		comandos.EjecutarEstado()

	default:

		fmt.Println(
			"Comando desconocido:",
			comando,
		)
	}
}
