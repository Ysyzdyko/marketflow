package httptransport

import (
	ports "marketflow/internal/ports/inbound"
	"net/http"
)

func RegisterRouters(svc ports.APIPorts, router *http.ServeMux) {
	h, err := NewHandler(svc)
	if err != nil {
		panic("failed to create handler:  " + err.Error())
	}

	router.HandleFunc("GET /prices/latest/{symbol}", h.GetLastPrice)
	router.HandleFunc("GET /prices/latest/{exchange}/{symbol}", h.GetLastPrice)
	router.HandleFunc("GET /prices/highest/{symbol}", h.GetHighestPrice)
	router.HandleFunc("GET /prices/highest/{exchange}/{symbol}", h.GetHighestPrice)
	router.HandleFunc("GET /prices/highest/{symbol}?period={duration}", h.GetHighestPrice)
	router.HandleFunc("GET /prices/highest/{exchange}/{symbol}?period={duration}", h.GetHighestPrice)
	router.HandleFunc("GET /prices/lowest/{symbol}", h.GetLowestPrice)
	router.HandleFunc("GET /prices/lowest/{exchange}/{symbol}", h.GetLowestPrice)
	router.HandleFunc("GET /prices/lowest/{symbol}?period={duration}", h.GetLowestPrice)
	router.HandleFunc("GET /prices/lowest/{exchange}/{symbol}?period={duration}", h.GetLowestPrice)
	router.HandleFunc("GET /prices/average/{symbol}", h.GetAvgPrice)
	router.HandleFunc("GET /prices/average/{exchange}/{symbol}", h.GetAvgPrice)
	router.HandleFunc("GET /prices/average/{exchange}/{symbol}?period={duration}", h.GetAvgPrice)
	router.HandleFunc("GET /health", h.GetHealth)
	router.HandleFunc("POST /mode/{mode}", h.PostModeSwitcher)
}
