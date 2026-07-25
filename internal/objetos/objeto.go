package objetos

type TipoObjeto string

const (
	TipoBlob   TipoObjeto = "blob"
	TipoTree   TipoObjeto = "tree"
	TipoCommit TipoObjeto = "commit"
)

type Objeto struct {
	Tipo TipoObjeto

	Contenido []byte

	Hash string
}
