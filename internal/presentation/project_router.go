package presentation

import "net/http"

func ProjectRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		RenderJson(w, map[string]string{"project": "x"}, 202)
	})

	return router
}
