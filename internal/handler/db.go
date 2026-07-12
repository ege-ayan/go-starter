package handler

import (
	"context"
	"net/http"
	"time"
)

type dbInfoResponse struct {
	Connected bool     `json:"connected"`
	Version   string   `json:"version,omitempty"`
	Greetings []string `json:"greetings,omitempty"`
}

func (h *Handler) DBInfo(w http.ResponseWriter, r *http.Request) {
	if h.db == nil {
		writeJSON(w, http.StatusServiceUnavailable, dbInfoResponse{Connected: false})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	version, err := h.db.Version(ctx)
	if err != nil {
		h.logger.Error("database version query failed", "error", err)
		writeJSON(w, http.StatusServiceUnavailable, dbInfoResponse{Connected: false})
		return
	}

	greetings, err := h.db.ListGreetings(ctx)
	if err != nil {
		h.logger.Error("database greetings query failed", "error", err)
		writeJSON(w, http.StatusServiceUnavailable, dbInfoResponse{Connected: false})
		return
	}

	writeJSON(w, http.StatusOK, dbInfoResponse{
		Connected: true,
		Version:   version,
		Greetings: greetings,
	})
}
