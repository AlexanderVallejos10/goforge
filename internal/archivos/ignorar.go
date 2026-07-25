package archivos

import (
"path/filepath"
)


func DebeIgnorar(
ruta string,
) bool {


nombre := filepath.Base(
ruta,
)


if nombre == ".git" {
return true
}


if nombre == ".goforge" {
return true
}


return false

}
