package presentation

import (
	"beholder/internal/boundaries"
	"beholder/internal/data"
	"beholder/internal/presentation/controllers"
	"net/http"

	"github.com/uptrace/bun"
)

func UserRouter(db *bun.DB) http.Handler {
	router := http.NewServeMux()
	controller := controllers.NewUserController(data.NewUserRepository(db))
	router.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		input, err := ParseBody[boundaries.CreateUserInput](r.Body)
		if err != nil {
			RenderJson(w, err, 422)
			return
		}
		user, err := controller.CreateUser(*input)
		if err != nil {
			RenderJson(w, err, 400)
			return
		}
		RenderJson(w, user, 201)
	})

	return router
}
