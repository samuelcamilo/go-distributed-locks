package main

import (
	"os"

	"example.com/event-query-api/internal/controllers"
	"example.com/event-query-api/internal/core/handlers"
	"example.com/event-query-api/internal/repositories"
	"example.com/event-query-api/internal/services"
	"example.com/event-query-api/pkg/database"
	"example.com/event-query-api/pkg/logger"
	"example.com/event-query-api/pkg/server"
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
	controllers.Event.RegisterRouters(router)

	log.Info("starting server in port: ", os.Getenv("PORT"))
	if err := router.ListenAndServe(os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
