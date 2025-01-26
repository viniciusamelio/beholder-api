package main

import (
	"beholder-api/cmd/api/router"
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/services"

	_ "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	services.InitSomClient()

	fx.New(
		fx.Provide(
			services.NewSomDatasource,
			repositories.NewCallRepository,
			repositories.NewEnvironmentRepository,
			repositories.NewSessionRepository,
		),
		fx.Invoke(
			router.Router,
		),
	).Run()

}
