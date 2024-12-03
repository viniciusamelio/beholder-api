package main

import (
	db "beholder/internal/infra"
	"beholder/internal/presentation"
)

func main() {
	db.RunMigrations()
	db := db.GetDB()
	defer db.Close()

	presentation.SetupRouter()
}
