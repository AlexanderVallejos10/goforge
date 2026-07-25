package diferencias

import "strings"

type Linea struct {
	Tipo      string
	Contenido string
}

const (
	LineaIgual     = "igual"
	LineaEliminada = "eliminada"
	LineaAgregada  = "agregada"
)

func CompararLineas(
	contenidoAnterior []byte,
	contenidoActual []byte,
) []Linea {

	lineasAnteriores := dividirLineas(
		string(contenidoAnterior),
	)

	lineasActuales := dividirLineas(
		string(contenidoActual),
	)

	tabla := construirTablaLCS(
		lineasAnteriores,
		lineasActuales,
	)

	return reconstruirDiferencias(
		lineasAnteriores,
		lineasActuales,
		tabla,
	)
}

func dividirLineas(
	contenido string,
) []string {

	contenido = strings.ReplaceAll(
		contenido,
		"\r\n",
		"\n",
	)

	if contenido == "" {
		return []string{}
	}

	return strings.Split(
		contenido,
		"\n",
	)
}

func construirTablaLCS(
	anteriores []string,
	actuales []string,
) [][]int {

	tabla := make(
		[][]int,
		len(anteriores)+1,
	)

	for i := range tabla {
		tabla[i] = make(
			[]int,
			len(actuales)+1,
		)
	}

	for i := len(anteriores) - 1; i >= 0; i-- {

		for j := len(actuales) - 1; j >= 0; j-- {

			if anteriores[i] == actuales[j] {

				tabla[i][j] =
					tabla[i+1][j+1] + 1

				continue
			}

			if tabla[i+1][j] >=
				tabla[i][j+1] {

				tabla[i][j] =
					tabla[i+1][j]

			} else {

				tabla[i][j] =
					tabla[i][j+1]
			}
		}
	}

	return tabla
}

func reconstruirDiferencias(
	anteriores []string,
	actuales []string,
	tabla [][]int,
) []Linea {

	var resultado []Linea

	i := 0
	j := 0

	for i < len(anteriores) &&
		j < len(actuales) {

		if anteriores[i] == actuales[j] {

			resultado = append(
				resultado,
				Linea{
					Tipo:      LineaIgual,
					Contenido: anteriores[i],
				},
			)

			i++
			j++

			continue
		}

		if tabla[i+1][j] >=
			tabla[i][j+1] {

			resultado = append(
				resultado,
				Linea{
					Tipo:      LineaEliminada,
					Contenido: anteriores[i],
				},
			)

			i++

		} else {

			resultado = append(
				resultado,
				Linea{
					Tipo:      LineaAgregada,
					Contenido: actuales[j],
				},
			)

			j++
		}
	}

	for i < len(anteriores) {

		resultado = append(
			resultado,
			Linea{
				Tipo:      LineaEliminada,
				Contenido: anteriores[i],
			},
		)

		i++
	}

	for j < len(actuales) {

		resultado = append(
			resultado,
			Linea{
				Tipo:      LineaAgregada,
				Contenido: actuales[j],
			},
		)

		j++
	}

	return resultado
}
