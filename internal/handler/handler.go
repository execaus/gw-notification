package handler

import (
	"context"
	"gw-notification/config"
	"gw-notification/internal/domain"
	"gw-notification/internal/service"
	"strconv"
	"time"

	gweventbus "github.com/execaus/gw-event-bus"
	"github.com/execaus/gw-event-bus/message"
	"go.uber.org/zap"
)

type Handler struct {
	s        service.Service
	consumer gweventbus.Consumer
}

func (h *Handler) handlePaymentsHighValueTransfer(message message.PaymentsHighValueTransferMessage) {
	if err := h.s.Save(context.Background(), domain.Exchange{
		Email:     message.Email,
		From:      message.From,
		To:        message.To,
		Amount:    strconv.FormatFloat(float64(message.Amount), 'f', -1, 32),
		CreatedAt: time.Now(),
	}); err != nil {
		zap.L().Error(err.Error())
	}
}

func NewHandler(ctx context.Context, s service.Service, cfg config.EventBusConfig) *Handler {
	h := &Handler{
		s:        s,
		consumer: gweventbus.NewConsumer(cfg.Host, cfg.Port, zap.L()),
	}

	h.consumer.Topics.PaymentsHighValueTransfer.Handle(ctx, h.handlePaymentsHighValueTransfer)

	return h
}

func (h *Handler) Close() error {
	return h.consumer.Close()
}
