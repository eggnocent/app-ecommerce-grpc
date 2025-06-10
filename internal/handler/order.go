package handler

import (
	"context"
	"github/eggnocent/app-grpc-eccomerce/internal/service"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"github/eggnocent/app-grpc-eccomerce/pb/order"
)

type orderHandler struct {
	order.UnimplementedOrderServiceServer

	orderService service.IOrderService
}

func (oh *orderHandler) CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {

	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &order.CreateOrderResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	req, err := oh.orderService.CreateOrder(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (oh *orderHandler) ListOrderAdmin(ctx context.Context, request *order.ListOrderAdminRequest) (*order.ListOrderAdminResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &order.ListOrderAdminResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	req, err := oh.orderService.ListOrderAdmin(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (oh *orderHandler) ListOrder(ctx context.Context, request *order.ListOrderRequest) (*order.ListOrderResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &order.ListOrderResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	req, err := oh.orderService.ListOrder(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (oh *orderHandler) DetailOrder(ctx context.Context, request *order.DetailOrderRequest) (*order.DetailOrderResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &order.DetailOrderResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	req, err := oh.orderService.DetailOrder(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (oh *orderHandler) UpdateOrderStatus(ctx context.Context, request *order.UpdateOrderStatusRequest) (*order.UpdateOrderStatusResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &order.UpdateOrderStatusResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	req, err := oh.orderService.UpdateOrderStatus(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewOrderHandler(orderService service.IOrderService) *orderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}
