package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/swanden/rentateam/internal/controller/grpc"
	v1 "github.com/swanden/rentateam/internal/controller/http/v1"
	"github.com/swanden/rentateam/internal/domain/post/usecase"
	"github.com/swanden/rentateam/internal/infrastructure/domain/post/repository"
	"github.com/swanden/rentateam/pkg/config"
	"github.com/swanden/rentateam/pkg/grpcserver"
	"github.com/swanden/rentateam/pkg/httpserver"
	"github.com/swanden/rentateam/pkg/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(configFile string) {
	conf, err := config.New(configFile)
	if err != nil {
		log.Fatalf("Config error: %s\n", err)
	}

	pg, err := postgres.New(conf.Postgres.DSN, postgres.MaxPoolSize(conf.Postgres.MaxPoolSize))
	if err != nil {
		log.Fatalf("PostgreSQL error: %s\n", err)

	}
	defer pg.Close()

	repo := repository.New(pg)
	useCase := usecase.New(repo)
	validate := validator.New()

	handler := gin.New()
	v1.NewRouter(handler, repo, useCase, validate)
	httpServer := httpserver.New(handler, httpserver.Port(conf.HTTP.Port))

	fmt.Println("app - Run - http server started on http://localhost:" + conf.HTTP.Port)

	grpcServer, err := grpcserver.New(grpc.NewPost(repo, useCase), grpcserver.Port(conf.GRPC.Port))
	if err != nil {
		log.Fatalf("grpc server error: %s\n", err)
	}

	fmt.Println("app - Run - grpc server started on http://localhost:" + conf.GRPC.Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		fmt.Fprintf(os.Stderr, "app - Run - signal: %s\n", s.String())
	case err = <-httpServer.Notify():
		fmt.Fprintf(os.Stderr, "app - Run - signal: %s\n", err)
	case err = <-grpcServer.Notify():
		fmt.Fprintf(os.Stderr, "app - Run - signal: %s\n", err)
	}

	err = httpServer.Shutdown()
	if err != nil {
		fmt.Fprintf(os.Stderr, "app - Run - signal: %s", err)
	}

	grpcServer.Shutdown()
}
