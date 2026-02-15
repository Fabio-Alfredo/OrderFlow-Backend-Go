package bootstrap

import (
	"Auth-Service/internal/controller"
	"Auth-Service/internal/database"
	"Auth-Service/internal/repository/user"
	"Auth-Service/internal/service/auth"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Init() {

	configs, err := config.Load(".")
	if err != nil {
		panic(err)
	}

	port := configs.GetString("server.port")
	basePath := configs.GetString("server.base-path")
	registerPath := configs.GetString("server.register-path")

	log := logger.NewLogger()
	ctx := context.Background()

	router := http.NewServeMux()

	sqlDb := database.NewSQLConfig(configs)
	db, err := sqlDb.GetDB()
	if err != nil {
		panic(err)
	}

	defer sqlDb.CloseDb()
	parsers := setupParser(configs)
	userRepository := user.NewUserRepository(configs, db, log, parsers)
	registerService := auth.NewAuthService(configs, log, userRepository, parsers)
	registerController := controller.NewRegisterController(log, registerService, parsers)
	router.HandleFunc(basePath+registerPath, registerController.Controller)

	signals := make(chan os.Signal)
	errors := make(chan error)

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", port),
	}

	go func() {
		log.Info(ctx, "Service connections on port", "PORT", port)
		errors <- server.ListenAndServe()
	}()

	signal.Notify(signals, syscall.SIGINT)
	signal.Notify(signals, syscall.SIGTERM)

	select {
	case s := <-signals:
		log.Info(ctx, "Signal received ", "signal", s.String())
		break
	case e := <-errors:
		log.Error(ctx, "Error occurred", "error", e.Error())
		break
	}

}
