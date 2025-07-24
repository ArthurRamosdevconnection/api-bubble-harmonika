package teste

import "github.com/ArthurRamosdevconnection/api-bubble-harmonika/crud"

type TesteModel struct {
	crud.BaseModel
	Description string `json:"description"`
	Value       string `json:"value"`
	Date        string `json:"date"`
}
