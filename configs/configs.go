package configs

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func Load(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
	}
	return os.Getenv(key)
}

func FiberConfig() fiber.Config {
	return fiber.Config{
		StreamRequestBody: true,
		ReadTimeout:       90 * time.Second,
		WriteTimeout:      90 * time.Second,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
	}
}
