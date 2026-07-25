package objetos

type TipoObjeto string

const (
	TipoBlob TipoObjeto = "blob"
)

type Objeto struct {
	Tipo TipoObjeto

	Contenido []byte

	Hash string
}
