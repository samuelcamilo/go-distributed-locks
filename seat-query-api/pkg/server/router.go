package server

import (
	"net/http"
)

type (
	IMuxRouter interface {
		Server(port string) http.Server
		Get(path string, f HandlerFunc)
	}
	HandlerFunc func(ctx IMuxContext)
	serveMux    struct {
		router *http.ServeMux
	}
)

func NewMuxRouter() IMuxRouter {
	router := http.NewServeMux()

	return &serveMux{
		router: router,
	}
}

func (s *serveMux) Server(port string) http.Server {
	return http.Server{
		Addr:    port,
		Handler: s.router,
	}
}

func (s *serveMux) Get(path string, f HandlerFunc) {
	s.router.HandleFunc("GET "+path, func(w http.ResponseWriter, r *http.Request) {
		f(newMuxContext(w, r))
	})
}
