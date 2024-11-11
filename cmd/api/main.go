package main

import (
	"log"

	env "github.com/grvbrk/go-backend"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config{
		addr: env.GetEnv("ADDR", ":8000"),
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.serve(mux))

}
