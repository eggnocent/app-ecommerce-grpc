package service

import (
	"context"
	"fmt"
	"github/eggnocent/app-grpc-eccomerce/internal/entity"
	jwtentity "github/eggnocent/app-grpc-eccomerce/internal/entity/jwt"
	"github/eggnocent/app-grpc-eccomerce/internal/repository"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"github/eggnocent/app-grpc-eccomerce/pb/product"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type IProductService interface {
	CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreatedProductResponse, error)
	DetailProduct(ctx context.Context, request *product.DetailProductRequest) (*product.DetailProductResponse, error)
	EditProduct(ctx context.Context, request *product.EditProductRequest) (*product.EditProductResponse, error)
	DeleteProduct(ctx context.Context, request *product.DeleteProductRequest) (*product.DeleteProductResponse, error)
	ListProduct(ctx context.Context, request *product.ListProductRequest) (*product.ListProductResponse, error)
	ListProductAdmin(ctx context.Context, request *product.ListProductAdminRequest) (*product.ListProductAdminResponse, error)
	HighlightProducts(ctx context.Context, request *product.HighlightProductsRequest) (*product.HighlightProductsResponse, error)
}

type productService struct {
	productRespository repository.IProductRepository
}

func (ps *productService) CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreatedProductResponse, error) {

	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != entity.UserRoleAdmin {
		return nil, utils.UnauthenticatedResponse()
	}

	imagePath := filepath.Join("storage", "product", request.ImageFileName)
	_, err = os.Stat(imagePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &product.CreatedProductResponse{
				Base: utils.BadRequestResponse("file not found"),
			}, nil
		}
		return nil, err
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

func (ps *productService) DetailProduct(ctx context.Context, request *product.DetailProductRequest) (*product.DetailProductResponse, error) {
	productEntity, err := ps.productRespository.GetProductByID(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	if productEntity == nil {
		return &product.DetailProductResponse{
			Base: utils.NotFoundResponse("product not found"),
		}, nil
	}

	return &product.DetailProductResponse{
		Base:        utils.SuccessResponse("get product detail success"),
		Id:          productEntity.Id,
		Name:        productEntity.Name,
		Description: productEntity.Description,
		Price:       productEntity.Price,
		ImageUrl:    fmt.Sprintf("%s/product/%s", os.Getenv("STORAGE_SERVICE_URL"), productEntity.ImageFileName),
	}, nil
}
func (ps *productService) EditProduct(ctx context.Context, request *product.EditProductRequest) (*product.EditProductResponse, error) {
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != entity.UserRoleAdmin {
		return nil, utils.UnauthenticatedResponse()
	}

	productEntity, err := ps.productRespository.GetProductByID(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	if productEntity == nil {
		return &product.EditProductResponse{
			Base: utils.NotFoundResponse("product failed to update"),
		}, nil
	}

	if productEntity.ImageFileName != request.ImageFileName {
		newImagePath := filepath.Join("storage", "product", request.ImageFileName)
		_, err := os.Stat(newImagePath)
		if err != nil {
			if os.IsNotExist(err) {
				return &product.EditProductResponse{
					Base: utils.BadRequestResponse("image not found"),
				}, nil
			}

			return nil, err
		}

		oldImagePath := filepath.Join("storage", "product", productEntity.ImageFileName)
		err = os.Remove(oldImagePath)
		if err != nil {
			return nil, err
		}
	}

	newProduct := entity.Product{
		Id:            request.Id,
		Name:          request.Name,
		Description:   request.Description,
		Price:         request.Price,
		ImageFileName: request.ImageFileName,
		UpdatedAt:     time.Now(),
		UpdatedBy:     &claims.FullName,
	}
	err = ps.productRespository.UpdateProduct(ctx, &newProduct)
	if err != nil {
		return nil, err
	}

	return &product.EditProductResponse{
		Base: utils.SuccessResponse("edit product success"),
		Id:   request.Id,
	}, nil
}

func (ps *productService) DeleteProduct(ctx context.Context, request *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != entity.UserRoleAdmin {
		return nil, utils.UnauthenticatedResponse()
	}

	productEntity, err := ps.productRespository.GetProductByID(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	if productEntity == nil {
		return &product.DeleteProductResponse{
			Base: utils.NotFoundResponse("product failed to update"),
		}, nil
	}

	err = ps.productRespository.DeleteProduct(ctx, request.Id, time.Now(), claims.FullName)

	if err != nil {
		return nil, err
	}

	imagePath := filepath.Join("storage", "product", productEntity.ImageFileName)
	if err := os.Remove(imagePath); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	return &product.DeleteProductResponse{
		Base: utils.SuccessResponse("delete product success"),
	}, nil
}

func (ps *productService) ListProduct(ctx context.Context, request *product.ListProductRequest) (*product.ListProductResponse, error) {

	products, paginationResponse, err := ps.productRespository.GetProductPagination(ctx, request.Pagination)
	if err != nil {
		return nil, err
	}

	var data []*product.ListProductResponseItem = make([]*product.ListProductResponseItem, 0)
	for _, prod := range products {
		data = append(data, &product.ListProductResponseItem{
			Id:          prod.Id,
			Name:        prod.Name,
			Description: prod.Description,
			Price:       prod.Price,
			ImageUrl:    fmt.Sprintf("%s/product/%s", os.Getenv("STORAGE_SERVICE_URL"), prod.ImageFileName),
		})
	}

	return &product.ListProductResponse{
		Base:       utils.SuccessResponse("get list product success"),
		Pagination: paginationResponse,
		Data:       data,
	}, nil
}

func (ps *productService) ListProductAdmin(ctx context.Context, request *product.ListProductAdminRequest) (*product.ListProductAdminResponse, error) {
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != entity.UserRoleAdmin {
		return nil, utils.UnauthenticatedResponse()
	}

	products, paginationResponse, err := ps.productRespository.GetProductPaginationAdmin(ctx, request.Pagination)
	if err != nil {
		return nil, err
	}

	data := make([]*product.ListProductAdminResponseItem, 0, len(products))
	for _, prod := range products {
		data = append(data, &product.ListProductAdminResponseItem{
			Id:          prod.Id,
			Name:        prod.Name,
			Description: prod.Description,
			Price:       prod.Price,
			ImageUrl:    fmt.Sprintf("%s/product/%s", os.Getenv("STORAGE_SERVICE_URL"), prod.ImageFileName),
		})
	}

	return &product.ListProductAdminResponse{
		Base:       utils.SuccessResponse("get list product admin success"),
		Pagination: paginationResponse,
		Data:       data,
	}, nil
}

func (ps *productService) HighlightProducts(ctx context.Context, request *product.HighlightProductsRequest) (*product.HighlightProductsResponse, error) {
	products, err := ps.productRespository.GetProductHighlight(ctx)
	if err != nil {
		return nil, err
	}

	data := make([]*product.HighlightProductsResponseItem, 0, len(products))
	for _, prod := range products {
		data = append(data, &product.HighlightProductsResponseItem{
			Id:          prod.Id,
			Name:        prod.Name,
			Description: prod.Description,
			Price:       prod.Price,
			ImageUrl:    fmt.Sprintf("%s/product/%s", os.Getenv("STORAGE_SERVICE_URL"), prod.ImageFileName),
		})
	}

	return &product.HighlightProductsResponse{
		Base: utils.SuccessResponse("get highlight products success"),
		Data: data,
	}, nil
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &productService{
		productRespository: productRepository,
	}
}
