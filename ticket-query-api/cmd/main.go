package main

import (
	"os"

	"example.com/ticket-query-api/internal/controllers"
	"example.com/ticket-query-api/internal/core/handlers"
	"example.com/ticket-query-api/internal/repositories"
	"example.com/ticket-query-api/internal/services"
	"example.com/ticket-query-api/pkg/database"
	"example.com/ticket-query-api/pkg/logger"
	"example.com/ticket-query-api/pkg/server"
)

func main() {
	var (
		log          = logger.NewLogrusLogger()
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
	controllers.Ticket.RegisterRouters(router)

	log.Info("starting server in port: ", os.Getenv("PORT"))
	if err := router.ListenAndServe(os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
