package seat

import (
	"net/http"

	"example.com/seat-query-api/internal/services"
	"example.com/seat-query-api/pkg/logger"
)

type (
	IHandler interface {
		FindById(w http.ResponseWriter, r *http.Request)
	}

	handlers struct {
		srv *services.Container
		log logger.Logger
	}
)

func New(srv *services.Container, log logger.Logger) IHandler {
	return &handlers{srv: srv, log: log}
}

func (hdl *handlers) FindById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	hdl.log.Info("PathValue: ", id)
}
