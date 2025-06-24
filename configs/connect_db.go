package configs

import (
	"fmt"
	"log"
	"strconv"
)

// ConnectDB conecta ao banco de dados
func ConnectDB() string {
	var err error
	p := Load("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println("Erro ao converter a porta do banco de dados")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable search_path=%s",
		Load("DB_HOST"), Load("DB_USER"), Load("DB_PASSWORD"), Load("DB_NAME"), port, Load("SCHEMA"))
	fmt.Println(dsn)
	return dsn
}
