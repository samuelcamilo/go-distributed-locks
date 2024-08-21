package seats

import (
	"example.com/seat-query-api/internal/services"
	"example.com/seat-query-api/pkg/logger"
	"example.com/seat-query-api/pkg/server"
)

type (
	IHandler interface {
		GetById(c server.IMuxContext)
	}

	handlers struct {
		log logger.Logger
		srv *services.Container
	}
)

func New(log logger.Logger, srv *services.Container) IHandler {
	return &handlers{log: log, srv: srv}
}

func (h *handlers) GetById(c server.IMuxContext) {
	id, err := c.PathIntValue("id")
	if err != nil {
		c.BadRequest(err)
		return
	}

	seats, err := h.srv.Seat.GetById(id)
	if err != nil {
		h.log.Error("error when try to get seats")
	}

	c.Ok(seats)
}
