package events

import (
	"example.com/event-query-api/internal/services"
	"example.com/event-query-api/pkg/logger"
	"example.com/event-query-api/pkg/server"
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
		h.log.Error(err)
		c.BadRequest(err)
		return
	}

	event, err := h.srv.Event.GetById(id)
	if err != nil {
		h.log.Error(err)
		responseHandler(c, err)
		return
	}

	c.OK(event)
}
