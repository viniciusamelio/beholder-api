package main

import (
	"beholder/internal/data"
	db "beholder/internal/infra"
	"context"
)

func main() {
	db.RunMigrations()
	db := db.GetDB()
	defer db.Close()

	org := new(data.OrganizationModel)
	db.NewSelect().Model(org).Where("id = ? ", 1).Scan(context.Background())
}
