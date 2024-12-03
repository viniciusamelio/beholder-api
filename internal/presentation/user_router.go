package presentation

import "net/http"

func UserRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]any)
		data["hello"] = "world"
		RenderJson(w, data, 200)
	})

	return router
}
