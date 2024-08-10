package seats

import (
	"net/http"

	"example.com/seat-query-api/internal/core/handlers"
)

type (
	IController interface {
		RegisterRouters(r *http.ServeMux)
	}
	controllers struct {
		hdl *handlers.Container
	}
)

func New(hdl *handlers.Container) IController {
	return &controllers{hdl: hdl}
}

func (ctrl *controllers) RegisterRouters(r *http.ServeMux) {
	r.HandleFunc("GET /seat/{id}", ctrl.hdl.Seat.FindById)
}
