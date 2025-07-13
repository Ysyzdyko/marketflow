package httptransport

import (
	ports "marketflow/internal/ports/inbound"
	"net/http"
)

type Server struct {
	addr   string
	router *http.ServeMux
	svc    ports.APIPorts
}

func NewHTTPServer(svc ports.APIPorts) *Server {
	router := newRouter(svc)

	addr := ":8080"

	return &Server{
		addr:   addr,
		router: router,
		svc:    svc,
	}
}

func newRouter(svc ports.APIPorts) *http.ServeMux {
	router := http.NewServeMux()
	RegisterRouters(svc, router)
	return router
}

func (s *Server) Serve() error {
	svc := &http.Server{
		Addr: s.addr,
	}
	return svc.ListenAndServe()
}
