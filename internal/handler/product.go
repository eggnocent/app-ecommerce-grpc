package handler

import (
	"context"
	"github/eggnocent/app-grpc-eccomerce/internal/service"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"github/eggnocent/app-grpc-eccomerce/pb/product"
)

type productHandler struct {
	product.UnimplementedProductServiceServer

	productService service.IProductService
}

func (ph *productHandler) CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreatedProductResponse, error) {

	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &product.CreatedProductResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	req, err := ph.productService.CreateProduct(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (ph *productHandler) DetailProduct(ctx context.Context, request *product.DetailProductRequest) (*product.DetailProductResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &product.DetailProductResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	req, err := ph.productService.DetailProduct(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (ph *productHandler) EditProduct(ctx context.Context, request *product.EditProductRequest) (*product.EditProductResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &product.EditProductResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	req, err := ph.productService.EditProduct(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (ph *productHandler) DeleteProduct(ctx context.Context, request *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &product.DeleteProductResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	req, err := ph.productService.DeleteProduct(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (ph *productHandler) ListProduct(ctx context.Context, request *product.ListProductRequest) (*product.ListProductResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &product.ListProductResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	req, err := ph.productService.ListProduct(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (ph *productHandler) ListProductAdmin(ctx context.Context, request *product.ListProductAdminRequest) (*product.ListProductAdminResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &product.ListProductAdminResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	req, err := ph.productService.ListProductAdmin(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (ph *productHandler) HighlightProducts(ctx context.Context, request *product.HighlightProductsRequest) (*product.HighlightProductsResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &product.HighlightProductsResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	req, err := ph.productService.HighlightProducts(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewProductHandler(productService service.IProductService) *productHandler {
	return &productHandler{
		productService: productService,
	}
}
