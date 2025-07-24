package configs

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/teste"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	isMigrating := Load("MIGRATE")
	if isMigrating != "true" {
		return
	}
	err := db.AutoMigrate(
		&teste.TesteModel{},
	)
	if err != nil {
		log.Fatalf("erro no migrar banco: %v", err.Error())
	}
	return
}
