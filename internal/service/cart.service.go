package service

import (
	"context"
	"errors"
	"fmt"
	"github/eggnocent/app-grpc-eccomerce/internal/entity"
	jwtentity "github/eggnocent/app-grpc-eccomerce/internal/entity/jwt"
	"github/eggnocent/app-grpc-eccomerce/internal/repository"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"github/eggnocent/app-grpc-eccomerce/pb/cart"
	"os"
	"time"

	"github.com/google/uuid"
)

type cartService struct {
	productRepository repository.IProductRepository
	cartRepository    repository.ICartRepository
}

type ICartService interface {
	AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error)
	ListCart(ctx context.Context, request *cart.ListCartRequest) (*cart.ListCartResponse, error)
	DeleteCart(ctx context.Context, request *cart.DeleteCartRequest) (*cart.DeleteCartResponse, error)
	UpdateQuantity(ctx context.Context, request *cart.UpdateCartQuantityRequest) (*cart.UpdateCartQuantityResponse, error)
}

func (cs *cartService) AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error) {
	// cek dahulu apakah product id ada di db?
	// cek ke db apakah product udah ada di cart user?
	// kalo udah ada -> update cart
	// kalo belum -> insert cart baru
	// response

	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	productEntity, err := cs.productRepository.GetProductByID(ctx, request.ProductId)
	if err != nil {
		return nil, err
	}

	if productEntity == nil {
		return &cart.AddProductToCartResponse{
			Base: utils.NotFoundResponse("product not found"),
		}, nil
	}

	cartEntity, err := cs.cartRepository.GetCartByProductAndUserID(ctx, request.ProductId, claims.Subject)
	if err != nil {
		return nil, err
	}

	if cartEntity != nil {
		now := time.Now()
		cartEntity.Quantity += 1
		cartEntity.UpdatedAt = &now
		cartEntity.UpdatedBy = &claims.Subject

		err = cs.cartRepository.UpdatedCart(ctx, cartEntity)

		if err != nil {
			return nil, err
		}
		return &cart.AddProductToCartResponse{
			Base: utils.SuccessResponse("add product to cart success"),
			Id:   cartEntity.ID,
		}, nil
	}

	newCartEntity := entity.UserCart{
		ID:        uuid.NewString(),
		UserID:    claims.Subject,
		ProductID: request.ProductId,
		Quantity:  1,
		CreatedAt: time.Now(),
		CreatedBy: claims.FullName,
	}
	err = cs.cartRepository.CreateNewCart(ctx, &newCartEntity)
	if err != nil {
		return nil, err
	}

	return &cart.AddProductToCartResponse{
		Base: utils.SuccessResponse("add product to cart success"),
		Id:   newCartEntity.ID,
	}, nil
}

func (cs *cartService) ListCart(ctx context.Context, request *cart.ListCartRequest) (*cart.ListCartResponse, error) {
	// ambil auth user dari token
	// query list cart (join)
	// build response
	// kirimkan response

	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	carts, err := cs.cartRepository.GetListCart(ctx, claims.Subject)
	if err != nil {
		return nil, err
	}

	var items []*cart.ListCartResponseItem

	for _, cartEntity := range carts {

		if cartEntity.Product == nil {
			return nil, errors.New("internal error: product data not found in cart")
		}

		item := cart.ListCartResponseItem{
			CartId:          cartEntity.ID,
			ProductId:       cartEntity.ProductID,
			ProductName:     cartEntity.Product.Name,
			ProductImageUrl: fmt.Sprintf("%s/product/%s", os.Getenv("STORAGE_SERVICE_URL"), cartEntity.Product.ImageFileName),
			ProductPrice:    cartEntity.Product.Price,
			Quantity:        int64(cartEntity.Quantity),
		}

		items = append(items, &item)
	}

	return &cart.ListCartResponse{
		Base:  utils.SuccessResponse("get list card is success"),
		Items: items,
	}, nil

}

func (cs *cartService) DeleteCart(ctx context.Context, request *cart.DeleteCartRequest) (*cart.DeleteCartResponse, error) {
	// dapat user id
	// dapat data card
	// cocokan data user id di cart dengan auth
	// delete dengan hard delete
	// kirim response

	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	cartEntity, err := cs.cartRepository.GetCartByID(ctx, request.CartId)
	if err != nil {
		return nil, err
	}

	if cartEntity == nil {
		return &cart.DeleteCartResponse{
			Base: utils.NotFoundResponse("not found cart"),
		}, nil
	}

	if cartEntity.UserID != claims.Subject {
		return &cart.DeleteCartResponse{
			Base: utils.BadRequestResponse("user_id is not match"),
		}, nil
	}

	err = cs.cartRepository.DeleteCart(ctx, request.CartId)
	if err != nil {
		return nil, err
	}

	return &cart.DeleteCartResponse{
		Base: utils.SuccessResponse("delete cart success"),
	}, nil
}

func (cs *cartService) UpdateQuantity(ctx context.Context, request *cart.UpdateCartQuantityRequest) (*cart.UpdateCartQuantityResponse, error) {
	// get data cart by id
	// cocokan user id
	// update new quantity
	// update ke db
	// success response

	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	cartEntity, err := cs.cartRepository.GetCartByID(ctx, request.CartId)
	if err != nil {
		return nil, err
	}

	if cartEntity == nil {
		return &cart.UpdateCartQuantityResponse{
			Base: utils.NotFoundResponse("not found cart"),
		}, nil
	}

	if cartEntity.UserID != claims.Subject {
		return &cart.UpdateCartQuantityResponse{
			Base: utils.BadRequestResponse("cart user id is not match"),
		}, nil
	}

	if request.NewQuantity == 0 {
		err = cs.cartRepository.DeleteCart(ctx, request.CartId)
		if err != nil {
			return nil, err
		}

		return &cart.UpdateCartQuantityResponse{
			Base: utils.SuccessResponse("update cart quantity success"),
		}, nil
	}

	now := time.Now()
	cartEntity.Quantity = int(request.NewQuantity)
	cartEntity.UpdatedAt = &now
	cartEntity.UpdatedBy = &claims.FullName

	err = cs.cartRepository.UpdatedCart(ctx, cartEntity)
	if err != nil {
		return nil, err
	}

	return &cart.UpdateCartQuantityResponse{
		Base: utils.SuccessResponse("update cart quantity success"),
	}, nil

}

func NewCartService(productRepository repository.IProductRepository, cartRepository repository.ICartRepository) ICartService {
	return &cartService{
		productRepository: productRepository,
		cartRepository:    cartRepository,
	}
}
