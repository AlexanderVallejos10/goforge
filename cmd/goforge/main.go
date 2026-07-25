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

	case "commit":

		if len(os.Args) < 4 || os.Args[2] != "-m" {

			fmt.Println(
				"Usa: goforge commit -m mensaje",
			)

			return
		}

		comandos.EjecutarCommit(
			os.Args[3],
		)

	case "log":

		comandos.EjecutarLog()

	case "show":

		if len(os.Args) < 3 {

			fmt.Println(
				"Usa: goforge show hash",
			)

			return
		}

		comandos.EjecutarShow(
			os.Args[2],
		)

	case "branch":

		if len(os.Args) >= 3 {

			comandos.EjecutarBranch(
				os.Args[2],
			)

		} else {

			comandos.EjecutarBranch(
				"",
			)
		}

	case "checkout":

		if len(os.Args) < 3 {

			fmt.Println(
				"Usa: goforge checkout rama",
			)

			return
		}

		if os.Args[2] == "-b" {

			if len(os.Args) < 4 {

				fmt.Println(
					"Usa: goforge checkout -b rama",
				)

				return
			}

			comandos.EjecutarCheckout(
				true,
				os.Args[3],
			)

			return
		}

		comandos.EjecutarCheckout(
			false,
			os.Args[2],
		)

	case "diff":

		archivo := ""

		if len(os.Args) >= 3 {
			archivo = os.Args[2]
		}

		comandos.EjecutarDiff(
			archivo,
		)

	case "restore":

		if len(os.Args) < 3 {

			fmt.Println(
				"Usa: goforge restore archivo",
			)

			return
		}

		comandos.EjecutarRestore(
			os.Args[2],
		)

	case "rm":

		if len(os.Args) < 3 {

			fmt.Println(
				"Usa: goforge rm archivo",
			)

			return
		}

		comandos.EjecutarRm(
			os.Args[2],
		)

	case "reset":

		if len(os.Args) < 3 {

			fmt.Println(
				"Usa: goforge reset archivo",
			)

			return
		}

		comandos.EjecutarReset(
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
