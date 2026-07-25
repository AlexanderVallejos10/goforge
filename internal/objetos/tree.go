package objetos

type EntradaTree struct {
	Nombre string `json:"nombre"`

	Hash string `json:"hash"`
}

type Tree struct {
	Entradas []EntradaTree `json:"entradas"`
}
