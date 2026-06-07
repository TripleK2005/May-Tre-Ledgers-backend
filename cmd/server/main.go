package main

import (
	"log"

	"may-tre-ledger-be/internal/core/config"
	"may-tre-ledger-be/internal/core/database"
	"may-tre-ledger-be/internal/router"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Configuration loaded")

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")

	defer db.Close()

	r := router.Setup(cfg, db)

	log.Printf("server started at :%s", cfg.Port)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
