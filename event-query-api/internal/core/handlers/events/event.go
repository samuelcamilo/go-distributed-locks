package events

import (
	"example.com/event-query-api/internal/services"
	"example.com/event-query-api/pkg/logger"
	"example.com/event-query-api/pkg/server"
)

type (
	IHandler interface {
		GetAll(ctx server.IMuxContext)
		GetById(ctx server.IMuxContext)
	}
	handlers struct {
		log logger.Logger
		srv *services.Container
	}
)

func New(log logger.Logger, srv *services.Container) IHandler {
	return &handlers{log: log, srv: srv}
}

func (h *handlers) GetAll(ctx server.IMuxContext) {
	events, err := h.srv.Event.GetAll()
	if err != nil {
		h.log.Error(err)
		responseHandler(ctx, err)
		return
	}

	ctx.OK(events)
}

func (h *handlers) GetById(ctx server.IMuxContext) {
	id, err := ctx.PathIntValue("id")
	if err != nil {
		h.log.Error(err)
		ctx.BadRequest(err)
		return
	}

	event, err := h.srv.Event.GetById(id)
	if err != nil {
		h.log.Error(err)
		responseHandler(ctx, err)
		return
	}

	ctx.OK(event)
}
