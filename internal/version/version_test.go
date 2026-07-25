package version

import "testing"

func TestInformacionProyecto(t *testing.T) {
	if NombreProyecto == "" {
		t.Fatal("el nombre del proyecto no puede estar vacío")
	}

	if NumeroVersion == "" {
		t.Fatal("la versión del proyecto no puede estar vacía")
	}
}
