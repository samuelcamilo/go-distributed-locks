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
	eventHandler struct {
		log logger.Logger
		srv *services.Container
	}
)

func New(log logger.Logger, srv *services.Container) IHandler {
	return &eventHandler{log: log, srv: srv}
}

func (h *eventHandler) GetAll(ctx server.IMuxContext) {
	events, err := h.srv.Event.GetAll()
	if err != nil {
		h.log.Error(err)
		responseHandler(ctx, err)
		return
	}

	ctx.OK(events)
}

func (h *eventHandler) GetById(ctx server.IMuxContext) {
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
