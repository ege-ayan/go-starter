package handler

import "net/http"

type greetResponse struct {
	Message string `json:"message"`
}

func (h *Handler) Greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	writeJSON(w, http.StatusOK, greetResponse{
		Message: "Hello, " + name + "!",
	})
}
