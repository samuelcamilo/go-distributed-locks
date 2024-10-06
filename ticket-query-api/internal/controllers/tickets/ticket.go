package tickets

import (
	"example.com/ticket-query-api/internal/core/handlers"
	"example.com/ticket-query-api/pkg/server"
)

type (
	IController interface {
		RegisterRouters(r server.IMuxRouter)
	}
	controllers struct {
		hdl *handlers.Container
	}
)

func New(hdl *handlers.Container) IController {
	return &controllers{hdl: hdl}
}

func (ctrl *controllers) RegisterRouters(r server.IMuxRouter) {
	r.Get("/tickets/", ctrl.hdl.Ticket.GetAll)
	r.Get("/tickets/{id}", ctrl.hdl.Ticket.GetById)
}
