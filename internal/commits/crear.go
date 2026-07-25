package commits

import (
	"encoding/json"

	"github.com/AlexanderVallejos10/goforge/internal/indice"
	"github.com/AlexanderVallejos10/goforge/internal/objetos"
)

func CrearTree(
	entradas []indice.Entrada,
) ([]byte, error) {

	var archivos []objetos.EntradaTree

	for _, entrada := range entradas {

		archivos = append(
			archivos,
			objetos.EntradaTree{
				Nombre: entrada.Archivo,
				Hash:   entrada.Hash,
			},
		)

	}

	tree := objetos.Tree{
		Entradas: archivos,
	}

	return json.Marshal(
		tree,
	)

}
