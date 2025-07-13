package httptransport

import (
	"fmt"
	"log/slog"
	ports "marketflow/internal/ports/inbound"
	"net/http"
)

type Handler struct {
	svc ports.APIPorts
	// mode model.AppMode
}

func NewHandler(svc ports.APIPorts) (*Handler, error) {
	return &Handler{
		svc: svc,
	}, nil
}

func (h *Handler) GetLastPrice(w http.ResponseWriter, r *http.Request) {
	//symbol
	//exchange/symbol
}

func (h *Handler) GetHighestPrice(w http.ResponseWriter, r *http.Request) {
	//symbol
	//exchange/symbol
	//symbol?period=duration
	//exchange/symbol?period=duration
}

func (h *Handler) GetLowestPrice(w http.ResponseWriter, r *http.Request) {
	//symbol
	//exchange/symbol
	//symbol?period=duration
	//exchange/symbol?period=duration
}

func (h *Handler) GetAvgPrice(w http.ResponseWriter, r *http.Request) {
	//symbol
	//exchange/symbol
	//exchange/symbol?period=duration
}

func (h *Handler) PostModeSwitcher(w http.ResponseWriter, r *http.Request) {
	m := r.PathValue("mode")

	switch m {
	case "test":
		h.svc.SetMode(false)
		slog.Info("App in TEST mode")
	case "live":
		h.svc.SetMode(true)
		slog.Info("App in LIVE mode")
	default:
		http.Error(w, "invalid mode", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Mode switched to %s", m)
}

func (h *Handler) GetHealth(w http.ResponseWriter, r *http.Request) {}
