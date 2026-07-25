package objetos

import (
	"crypto/sha256"
	"encoding/hex"
)

func CalcularHash(contenido []byte) string {

	hash := sha256.Sum256(contenido)

	return hex.EncodeToString(
		hash[:],
	)

}
