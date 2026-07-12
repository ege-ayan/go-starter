package handler

import "net/http"

type statusResponse struct {
	App     string `json:"app"`
	Version string `json:"version"`
	Env     string `json:"env"`
}

func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, statusResponse{
		App:     h.cfg.AppName,
		Version: h.cfg.AppVersion,
		Env:     h.cfg.Env,
	})
}
