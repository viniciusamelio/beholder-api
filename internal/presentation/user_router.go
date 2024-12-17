package presentation

import (
	"beholder/internal/boundaries"
	"beholder/internal/data"
	"beholder/internal/exception"
	"beholder/internal/presentation/controllers"
	"net/http"
	"strconv"

	"github.com/uptrace/bun"
)

func UserRouter(db *bun.DB) http.Handler {
	router := http.NewServeMux()
	controller := controllers.NewUserController(data.NewUserRepository(db))
	router.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		input, err := ParseBody[boundaries.CreateUserInput](r.Body)
		if err != nil {
			HandleException(w, &exception.InvalidArguments{Args: "Invalid body"})
			return
		}
		user, err := controller.CreateUser(*input)
		if err != nil {
			HandleException(w, &exception.ProcessingException{Args: "Creating user was not possible"})
			return
		}
		RenderJson(w, user, 201)
	})

	router.HandleFunc("DELETE /{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			HandleException(w, &exception.InvalidArguments{Args: ""})
			return
		}
		input := boundaries.DeleteUserInput{
			ID: parsedId,
		}
		result, err := controller.DeleteUser(input)
		if err != nil {
			HandleException(w, &exception.ProcessingException{Args: "Deleting user was not possible"})
			return
		}

		RenderJson(w, result, 200)
	})

	return router
}
