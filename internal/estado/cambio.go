package estado

type TipoCambio string

const (
	TipoNuevo      TipoCambio = "NUEVO"
	TipoModificado TipoCambio = "MODIFICADO"
	TipoEliminado  TipoCambio = "ELIMINADO"
)

type Cambio struct {
	Archivo string
	Tipo    TipoCambio
}
