package service

import (
	"context"
	"fmt"
	"github/eggnocent/app-grpc-eccomerce/internal/entity"
	jwtentity "github/eggnocent/app-grpc-eccomerce/internal/entity/jwt"
	"github/eggnocent/app-grpc-eccomerce/internal/repository"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"github/eggnocent/app-grpc-eccomerce/pb/product"
	"time"

	"github.com/google/uuid"
)

type IProductService interface {
	CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreatedProductResponse, error)
}

type productService struct {
	productRespository repository.IProductRepository
}

func (ps *productService) CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreatedProductResponse, error) {

	fmt.Println("token ditemukan")
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		fmt.Println("token ditemukan di if")
		return nil, err
	}

	if claims.Role != entity.UserRoleAdmin {
		return nil, utils.UnauthenticatedResponse()
	}

	// cek email apakah sudah ada?
	productEntity := entity.Product{
		Id:            uuid.NewString(),
		Name:          request.Name,
		Description:   request.Description,
		Price:         request.Price,
		ImageFileName: request.ImageFileName,
		CreatedAt:     time.Now(),
		CreatedBy:     &claims.FullName,
		UpdatedAt:     time.Now(),
		IsDeleted:     false,
	}
	err = ps.productRespository.CreateNewProduct(ctx, &productEntity)
	if err != nil {
		return nil, err
	}

	return &product.CreatedProductResponse{
		Base: utils.SuccessResponse("Product is Successfuly created"),
		Id:   productEntity.Id,
	}, nil
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &productService{
		productRespository: productRepository,
	}
}
