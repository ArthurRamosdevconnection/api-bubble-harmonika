package nested_example

import "github.com/ArthurRamosdevconnection/api-bubble-harmonika/crud"

type StructPai struct {
	crud.BaseModel
	Nome   string
	Filhos []StructFilho
}

type StructFilho struct {
	crud.BaseModel
	Nome string
}
