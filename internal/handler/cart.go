package handler

import (
	"context"
	"github/eggnocent/app-grpc-eccomerce/internal/service"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"github/eggnocent/app-grpc-eccomerce/pb/cart"
)

type cartHandler struct {
	cart.UnimplementedCartServiceServer

	cartService service.ICartService
}

func (ch *cartHandler) AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &cart.AddProductToCartResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	resp, err := ch.cartService.AddProductToCart(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ch *cartHandler) ListCart(ctx context.Context, request *cart.ListCartRequest) (*cart.ListCartResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &cart.ListCartResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	resp, err := ch.cartService.ListCart(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ch *cartHandler) DeleteCart(ctx context.Context, request *cart.DeleteCartRequest) (*cart.DeleteCartResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &cart.DeleteCartResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	resp, err := ch.cartService.DeleteCart(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ch *cartHandler) UpdateCartQuantity(ctx context.Context, request *cart.UpdateCartQuantityRequest) (*cart.UpdateCartQuantityResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &cart.UpdateCartQuantityResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	resp, err := ch.cartService.UpdateQuantity(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewCartHandler(cartService service.ICartService) *cartHandler {
	return &cartHandler{
		cartService: cartService,
	}
}
