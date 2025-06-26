package main

import (
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/src/configs"
	"github.com/ArthurRamosdevconnection/api-bubble-harmonika/src/teste"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := configs.ConnectDB()

	fiberConfig := configs.FiberConfig()
	app := fiber.New(fiberConfig)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	app.Use(recover.New())

	app.Use(cors.New())

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
	configs.Migrate(db)

	teste.RegisterRoutes(app, db)
	log.Fatal(app.Listen(configs.Load("PORT")))
}
