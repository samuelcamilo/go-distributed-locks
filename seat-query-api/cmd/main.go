package main

import (
	"os"

	"example.com/seat-query-api/internal/controllers"
	"example.com/seat-query-api/internal/core/handlers"
	"example.com/seat-query-api/internal/repositories"
	"example.com/seat-query-api/internal/services"
	"example.com/seat-query-api/pkg/database"
	"example.com/seat-query-api/pkg/logger"
	"example.com/seat-query-api/pkg/server"
)

func main() {
	log := logger.NewLogrusLogger()

	var (
		router       = server.NewMuxRouter()
		repositories = repositories.New(repositories.Options{
			Log:    log,
			Client: database.MustConnect(os.Getenv("MONGODB_URI")),
		})
		services = services.New(services.Options{
			Log:  log,
			Repo: repositories,
		})
		handlers = handlers.New(handlers.Options{
			Log: log,
			Srv: services,
		})
		controllers = controllers.New(controllers.Options{
			Hdl: handlers,
		})
	)

	log.Info("adding middlewares")

	log.Info("registering routes")
	controllers.Seat.RegisterRouters(router)

	log.Info("starting server in port: ", os.Getenv("PORT"))
	if err := router.ListenAndServe(os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
