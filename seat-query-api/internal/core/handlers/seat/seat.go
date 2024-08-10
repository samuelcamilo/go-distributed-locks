package seat

import (
	"net/http"

	"example.com/seat-query-api/internal/services"
	"example.com/seat-query-api/pkg/logger"
	"example.com/seat-query-api/pkg/server"
)

type (
	IHandler interface {
		GetAll(c server.MuxContext)
	}

	handlers struct {
		srv *services.Container
		log logger.Logger
	}
)

func New(srv *services.Container, log logger.Logger) IHandler {
	return &handlers{srv: srv, log: log}
}

func (h *handlers) GetAll(c server.MuxContext) {
	seats, err := h.srv.Seat.GetAll()
	if err != nil {
		h.log.Error("error when try to get seats")
	}

	c.JSON(http.StatusOK, seats)
}
