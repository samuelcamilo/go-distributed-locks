package server

import (
	"net/http"
)

type (
	IMuxRouter interface {
		ListenAndServe(port string) error
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

func (s *serveMux) ListenAndServe(port string) error {
	serve := http.Server{
		Addr:    port,
		Handler: s.router,
	}

	return serve.ListenAndServe()
}

// func (s *serveMux) Server(port string) http.Server {
//	return http.Server{
//		Addr:    os.Getenv("PORT"),
//		Handler: s.router,
//	}
// }

func (s *serveMux) Get(path string, f HandlerFunc) {
	s.router.HandleFunc("GET "+path, func(w http.ResponseWriter, r *http.Request) {
		f(newMuxContext(w, r))
	})
}
