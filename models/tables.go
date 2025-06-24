package models

import (
	"gorm.io/gorm"
	"time"
)

type Teste struct {
	gorm.Model
	Description string    `json:"description"`
	Value       string    `json:"value"`
	Date        time.Time `json:"date"`
}
