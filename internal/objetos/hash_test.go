package objetos

import "testing"

func TestCalcularHash(t *testing.T) {

	contenido := []byte(
		"Hola GoForge",
	)

	hash := CalcularHash(
		contenido,
	)

	if hash == "" {

		t.Fatal(
			"el hash no debe estar vacío",
		)

	}

	if len(hash) != 64 {

		t.Fatalf(
			"el hash debe tener 64 caracteres, tiene %d",
			len(hash),
		)

	}

}
