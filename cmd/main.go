package main

import (
	"context"
	"gw-notification/config"
	"gw-notification/internal/handler"
	"gw-notification/internal/repository"
	"gw-notification/internal/service"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	ctx := context.Background()
	cfg := config.LoadConfig()

	r, closeFn := repository.NewRepository(ctx, cfg.Database)
	s := service.NewExchangeService(r)

	h := handler.NewHandler(ctx, s, cfg.EventBus)

	zap.L().Info("service started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	zap.L().Info("shutting down server...")
	if err := closeFn(ctx); err != nil {
		zap.L().Error(err.Error())
	}
	if err := h.Close(); err != nil {
		zap.L().Error(err.Error())
	}
	zap.L().Info("server stopped")
}
