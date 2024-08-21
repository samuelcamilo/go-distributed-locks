package seats

import (
	"example.com/seat-query-api/internal/core/handlers"
	"example.com/seat-query-api/pkg/server"
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
	r.Get("/seats/sessions/{id}", ctrl.hdl.Seat.GetById)
}
