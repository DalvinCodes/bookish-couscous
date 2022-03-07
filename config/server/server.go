package server

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	strictRouting, _ = strconv.ParseBool(os.Getenv("SERVER_STRICT_ROUTING"))
	caseSensitive, _ = strconv.ParseBool(os.Getenv("CASE_SENSITIVE"))
	read_timeout, _  = strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	write_timeout, _ = strconv.Atoi(os.Getenv("READ_TIMEOUT"))
)

func StartApplication(env string) *fiber.App {
	err := Environment.Load(env)
	if err != nil {
		log.Fatalf("unable to start application: %v", err)
	}
	godotenv.Load()

	app := configureFiberApp()

	return app

}

func configureFiberApp() *fiber.App {
	app := *fiber.New(fiber.Config{
		StrictRouting: strictRouting,
		CaseSensitive: caseSensitive,
		ReadTimeout:   time.Duration(read_timeout) * time.Second,
		WriteTimeout:  time.Duration(write_timeout) * time.Second,
	})
	return &app
}
