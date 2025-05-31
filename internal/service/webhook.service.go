package service

import (
	"context"
	"errors"
	"github/eggnocent/app-grpc-eccomerce/internal/dto"
	"github/eggnocent/app-grpc-eccomerce/internal/entity"
	"github/eggnocent/app-grpc-eccomerce/internal/repository"
	"log"
	"time"
)

type IwebHookService interface {
	ReceiveInvoice(ctx context.Context, request *dto.XenditInvoiceRequest) error
}

type webHookService struct {
	orderRepository repository.IOrderRepository
}

func (ws *webHookService) ReceiveInvoice(ctx context.Context, request *dto.XenditInvoiceRequest) error {
	// find order di db

	// ganti atau update entity

	// update ke db

	orderEntity, err := ws.orderRepository.GetOrderByID(ctx, request.ExternalID)
	if err != nil {
		return err
	}

	if orderEntity == nil {
		return errors.New("order not found")
	}

	now := time.Now()
	updatedBY := "System"
	orderEntity.OrderStatusCode = entity.OrderStatusCodePaid
	orderEntity.UpdatedAt = &now
	orderEntity.UpdatedBy = &updatedBY
	orderEntity.XenditPaidAt = &now
	orderEntity.XenditPaymentChannel = &request.PaymentChannel
	orderEntity.XenditPaymentMethod = &request.PaymentMethod

	err = ws.orderRepository.UpdateOrder(ctx, orderEntity)
	if err != nil {
		return err
	}

	log.Printf("Received invoice webhook: %+v\n", request)

	return nil
}

func NewWebHookService(orderRepository repository.IOrderRepository) IwebHookService {
	return &webHookService{
		orderRepository: orderRepository,
	}
}
