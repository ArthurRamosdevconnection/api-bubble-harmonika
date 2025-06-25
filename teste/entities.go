package teste

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/crud"
	"time"
)

type Teste struct {
	Description string    `json:"description"`
	Value       string    `json:"value"`
	Date        time.Time `json:"date"`
}

type TesteModel struct {
	crud.BaseModel
	Teste
}
