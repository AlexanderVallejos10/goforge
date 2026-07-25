package objetos

func Leer(
	rutaRepositorio string,
	hash string,
) (ObjetoLeido, error) {

	datos, err := LeerObjetoCompleto(
		rutaRepositorio,
		hash,
	)

	if err != nil {
		return ObjetoLeido{}, err
	}

	tipo, contenido, err := LeerContenidoObjeto(
		datos,
	)

	if err != nil {
		return ObjetoLeido{}, err
	}

	return ObjetoLeido{

		Tipo: tipo,

		Contenido: contenido,
	}, nil

}
