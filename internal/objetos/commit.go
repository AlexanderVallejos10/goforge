package objetos

import "time"

type Commit struct {
	Tree string `json:"tree"`

	Padre string `json:"padre"`

	Autor string `json:"autor"`

	Mensaje string `json:"mensaje"`

	Fecha time.Time `json:"fecha"`
}
