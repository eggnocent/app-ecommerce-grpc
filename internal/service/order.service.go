package service

import (
	"context"
	"fmt"
	"github/eggnocent/app-grpc-eccomerce/internal/entity"
	jwtentity "github/eggnocent/app-grpc-eccomerce/internal/entity/jwt"
	"github/eggnocent/app-grpc-eccomerce/internal/repository"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"github/eggnocent/app-grpc-eccomerce/pb/order"
	"time"

	"github.com/google/uuid"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error)
}

type orderService struct {
	orderRepository   repository.IOrderRepository
	productRepository repository.IProductRepository
}

func (os *orderService) CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {

	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	numbering, err := os.orderRepository.GetNumbering(ctx, "order")
	if err != nil {
		return nil, err
	}

	// Ambil produk
	var productIDs = make([]string, len(request.Products))
	for i := range request.Products {
		productIDs[i] = request.Products[i].Id
	}

	products, err := os.productRepository.GetProductByIDs(ctx, productIDs)
	if err != nil {
		return nil, err
	}

	productMap := make(map[string]*entity.Product)
	for i := range products {
		productMap[products[i].Id] = products[i]
	}

	var total float64 = 0
	for _, p := range request.Products {
		if productMap[p.Id] == nil {
			return &order.CreateOrderResponse{
				Base: utils.NotFoundResponse(fmt.Sprintf("product %s not found", p.Id)),
			}, nil
		}
		total += productMap[p.Id].Price * float64(p.Quantity)
	}

	now := time.Now()
	expiredAt := now.Add(24 * time.Hour)

	orderEntity := entity.Order{
		ID:              uuid.NewString(),
		Number:          fmt.Sprintf("ORD-%d%08d", now.Year(), numbering.Numbering),
		UserID:          claims.Subject,
		OrderStatusCode: entity.OrderStatusCodeUnpaid,
		UserFullName:    request.FullName,
		Address:         request.Address,
		PhoneNumber:     request.PhoneNumber,
		Notes:           &request.Notes,
		Total:           total,
		ExpiredAt:       &expiredAt,
		CreatedAt:       now,
		CreatedBy:       claims.FullName,
	}

	err = os.orderRepository.CreateOrder(ctx, &orderEntity)
	if err != nil {
		return nil, err
	}

	for _, p := range request.Products {
		orderItem := entity.OrderItem{
			ID:                   uuid.NewString(),
			ProductID:            p.Id,
			ProductName:          productMap[p.Id].Name,
			ProductImageFileName: productMap[p.Id].ImageFileName,
			ProductPrice:         productMap[p.Id].Price,
			Quantity:             p.Quantity,
			OrderID:              orderEntity.ID,
			CreatedAt:            now,
			CreatedBy:            claims.FullName,
		}

		fmt.Println("[Service] Simpan order item:", orderItem.ProductName)

		err = os.orderRepository.CreateOrderItem(ctx, &orderItem)
		if err != nil {
			fmt.Println("[Service] Gagal simpan order item:", err)
			return nil, err
		}
	}

	numbering.Numbering++
	err = os.orderRepository.UpdateNumbering(ctx, numbering)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{
		Base: utils.SuccessResponse("create order success"),
		Id:   orderEntity.ID,
	}, nil
}

func NewOrderService(orderRepository repository.IOrderRepository, productRepository repository.IProductRepository) IOrderService {
	return &orderService{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}
