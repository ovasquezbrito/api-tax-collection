package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/ovasquezbrito/tax-collection/docs"

	_ "github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"

	baseapp "github.com/ovasquezbrito/tax-collection"
	"github.com/ovasquezbrito/tax-collection/pkg/handler"
	"github.com/ovasquezbrito/tax-collection/pkg/repository"
	"github.com/ovasquezbrito/tax-collection/pkg/service"
	"github.com/ovasquezbrito/tax-collection/token"
	"github.com/ovasquezbrito/tax-collection/util"
)

// @title Tag Service API
// @version 1.0
// @description A tag service API in Go using Gin framework.

// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization

// @host localhost:8080
// @BasePath /
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	config, err := util.LoadConfig(".")
	if err != nil {
		logrus.Fatalf("error iniciando configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     config.Host,
		Port:     config.Port,
		Username: config.Username,
		DBName:   config.DBName,
		SSLMode:  config.SSLMode,
		Password: config.Password,
	})

	fmt.Println(config.AccessTokenDuration)
	fmt.Println(config.TokenSymmetricKey)

	if err != nil {
		logrus.Fatalf("error iniciando base de datos: %s", err.Error())
	}

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		fmt.Errorf("cannot create token maker: %w", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, tokenMaker, config)
	handlers := handler.NewHandler(services, tokenMaker)

	svr := new(baseapp.Server)
	go func() {
		if err := svr.Run(config.PortServer, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Print("App Iniciada")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	if err := svr.Shutdown(context.Background()); err != nil {
		logrus.Errorf("ah ocurrido un error en eñl servidor shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("ah ocurrido un error con la conexión de la bases de dato: %s", err.Error())
	}
}
