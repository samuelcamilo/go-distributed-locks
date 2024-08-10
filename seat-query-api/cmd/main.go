package main

import (
	"example.com/seat-query-api/config"
	"example.com/seat-query-api/internal/controllers"
	"example.com/seat-query-api/internal/core/handlers"
	"example.com/seat-query-api/internal/services"
	"example.com/seat-query-api/pkg/logger"
	"example.com/seat-query-api/pkg/server"
)

func main() {
	log := logger.NewLogrusLogger()
	configs, err := config.LoadConfig("./cmd")
	if err != nil {
		log.Fatal("failed to read config: ", err)
		return
	}

	var (
		router   = server.NewMuxRouter()
		services = services.New(services.Options{
			Log: log,
		})
		handlers = handlers.New(handlers.Options{
			Srv: services,
			Log: log,
		})
		controllers = controllers.New(controllers.Options{
			Hdl: handlers,
		})
	)

	log.Info("registering routes")
	controllers.Seat.RegisterRouters(router)

	server := router.Server(configs.Port)

	log.Info("starting server in port: ", configs.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
