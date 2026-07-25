package main

import (
	"fmt"

	"github.com/AlexanderVallejos10/goforge/internal/version"
)

func main() {
	fmt.Printf("%s %s\n", version.NombreProyecto, version.NumeroVersion)
}
