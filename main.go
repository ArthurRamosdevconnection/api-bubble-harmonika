package main

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/configs"
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func main() {
	dsn := configs.ConnectDB()

	fiberConfig := configs.FiberConfig()
	app := fiber.New(fiberConfig)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "testes.", // aqui define explicitamente o schema
			SingularTable: false,
		},
	})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	app.Use(recover.New())

	app.Use(cors.New())

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
	configs.Migrate(db)

	routes.RegisterAllRoutes(app, db)
	log.Fatal(app.Listen(configs.Load("PORT")))
}
