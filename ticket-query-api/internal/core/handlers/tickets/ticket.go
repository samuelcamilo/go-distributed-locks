package tickets

import (
	"example.com/ticket-query-api/internal/services"
	"example.com/ticket-query-api/pkg/logger"
	"example.com/ticket-query-api/pkg/server"
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
	tickets, err := h.srv.Ticket.GetAll(ctx.Context())
	if err != nil {
		h.log.Error(err)
		responseHandler(ctx, err)
		return
	}

	ctx.OK(tickets)
}

func (h *handlers) GetById(ctx server.IMuxContext) {
	id, err := ctx.PathIntValue("id")
	if err != nil {
		h.log.Error(err)
		ctx.BadRequest(err)
		return
	}

	ticket, err := h.srv.Ticket.GetById(ctx.Context(), id)
	if err != nil {
		h.log.Error(err)
		responseHandler(ctx, err)
		return
	}

	ctx.OK(ticket)
}
