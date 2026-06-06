package main

import (
	"log"

	"may-tre-ledger-be/internal/config"
	"may-tre-ledger-be/internal/database"
	"may-tre-ledger-be/internal/router"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := router.Setup()

	log.Printf("server started at :%s", cfg.Port)

	r.Run(":" + cfg.Port)
}
