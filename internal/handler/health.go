package handler

import (
	"context"
	"net/http"
	"time"
)

type healthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	resp := healthResponse{
		Status:   "ok",
		Database: "skipped",
	}

	if h.db != nil {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		if err := h.db.Ping(ctx); err != nil {
			h.logger.Error("database ping failed", "error", err)
			writeJSON(w, http.StatusServiceUnavailable, healthResponse{
				Status:   "degraded",
				Database: "down",
			})
			return
		}

		resp.Database = "ok"
	}

	writeJSON(w, http.StatusOK, resp)
}
