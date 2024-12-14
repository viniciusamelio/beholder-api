package presentation

import (
	"beholder/internal/infra"
	"net/http"
)

func SetupRouter() {
	router := http.NewServeMux()
	apiRouter := http.NewServeMux()

	db := infra.GetDB()

	apiRouter.Handle("/user", UserRouter(db))
	apiRouter.Handle("/project", ProjectRouter())

	router.Handle("/api/", http.StripPrefix("/api", apiRouter))

	http.ListenAndServe(":8080", router)
}
