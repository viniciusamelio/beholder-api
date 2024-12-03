package presentation

import (
	"net/http"
)

func SetupRouter() {
	router := http.NewServeMux()
	apiRouter := http.NewServeMux()

	apiRouter.Handle("/user", UserRouter())
	apiRouter.Handle("/project", ProjectRouter())

	router.Handle("/api/", http.StripPrefix("/api", apiRouter))

	http.ListenAndServe(":8080", router)
}
