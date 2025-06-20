package service

import (
	"context"
	"database/sql"
	"fmt"
	"github/eggnocent/app-grpc-eccomerce/internal/entity"
	jwtentity "github/eggnocent/app-grpc-eccomerce/internal/entity/jwt"
	"github/eggnocent/app-grpc-eccomerce/internal/repository"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"github/eggnocent/app-grpc-eccomerce/pb/order"
	"log"
	ops "os"
	"runtime/debug"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error)
	ListOrderAdmin(ctx context.Context, request *order.ListOrderAdminRequest) (*order.ListOrderAdminResponse, error)
	ListOrder(ctx context.Context, request *order.ListOrderRequest) (*order.ListOrderResponse, error)
	DetailOrder(ctx context.Context, request *order.DetailOrderRequest) (*order.DetailOrderResponse, error)
	UpdateOrderStatus(ctx context.Context, request *order.UpdateOrderStatusRequest) (*order.UpdateOrderStatusResponse, error)
}

type orderService struct {
	db                *sql.DB
	orderRepository   repository.IOrderRepository
	productRepository repository.IProductRepository
}

func (os *orderService) CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := os.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if e := recover(); e != nil {
			_ = tx.Rollback()
			debug.PrintStack()
			panic(e)
		}
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	orderRepo := os.orderRepository.WithTransaction(tx)
	productRepo := os.productRepository.WithTransaction(tx)

	numbering, err := orderRepo.GetNumbering(ctx, "order")
	if err != nil {
		return nil, err
	}

	var productIDs = make([]string, len(request.Products))
	for i := range request.Products {
		productIDs[i] = request.Products[i].Id
	}

	products, err := productRepo.GetProductByIDs(ctx, productIDs)
	if err != nil {

		if pqErr, ok := err.(*pq.Error); ok {
			log.Printf("[Service] Detail Postgres Error: %s", pqErr.Message)
		}

		for _, p := range request.Products {
			return &order.CreateOrderResponse{
				Base: utils.NotFoundResponse(fmt.Sprintf("product %s not found", p.Id)),
			}, nil
		}
	}

	productMap := make(map[string]*entity.Product)
	for i := range products {
		productMap[products[i].Id] = products[i]
	}

	var total float64 = 0
	for _, p := range request.Products {
		if productMap[p.Id] == nil {
			_ = tx.Rollback()
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

	invoiceItems := make([]xendit.InvoiceItem, 0)
	for _, p := range request.Products {
		prod := productMap[p.Id]
		if prod != nil {
			invoiceItems = append(invoiceItems, xendit.InvoiceItem{
				Name:     prod.Name,
				Price:    prod.Price,
				Quantity: int(p.Quantity),
			})
		}
	}

	xenditInvoice, xenditErr := invoice.CreateWithContext(ctx, &invoice.CreateParams{
		ExternalID: orderEntity.ID,
		Amount:     total,
		Customer: xendit.InvoiceCustomer{
			GivenNames: claims.FullName,
		},
		Currency:           "IDR",
		SuccessRedirectURL: fmt.Sprintf("%s/checkout/%s/success", ops.Getenv("FRONTEND_BASE_URL"), orderEntity.ID),
		Items:              invoiceItems,
	})

	if xenditErr != nil {
		err = xenditErr
		return nil, err
	}

	orderEntity.XenditInvoiceID = &xenditInvoice.ID
	orderEntity.XenditInvoiceUrl = &xenditInvoice.InvoiceURL

	err = orderRepo.CreateOrder(ctx, &orderEntity)
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

		err = orderRepo.CreateOrderItem(ctx, &orderItem)
		if err != nil {
			return nil, err
		}
	}

	numbering.Numbering++
	err = orderRepo.UpdateNumbering(ctx, numbering)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{
		Base: utils.SuccessResponse("create order success"),
		Id:   orderEntity.ID,
	}, nil
}

func (os *orderService) ListOrderAdmin(ctx context.Context, request *order.ListOrderAdminRequest) (*order.ListOrderAdminResponse, error) {
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != entity.UserRoleAdmin {
		return nil, utils.UnauthenticatedResponse()
	}

	orders, metadata, err := os.orderRepository.GetListOrderAdminPagination(ctx, request.Pagination)
	if err != nil {
		return nil, err
	}

	items := make([]*order.ListOrderAdminResponseItem, 0)
	for _, o := range orders {

		products := make([]*order.ListOrderAdminResponseItemProduct, 0)
		for _, oi := range o.Items {
			products = append(products, &order.ListOrderAdminResponseItemProduct{
				Id:       oi.ProductID,
				Name:     oi.ProductName,
				Price:    oi.ProductPrice,
				Quantity: oi.Quantity,
			})
		}

		orderStatusCode := o.OrderStatusCode
		if o.OrderStatusCode == entity.OrderStatusCodeUnpaid && time.Now().After(*o.ExpiredAt) {
			orderStatusCode = entity.OrderStatusCodeExpired
		}

		items = append(items, &order.ListOrderAdminResponseItem{
			Id:         o.ID,
			Number:     o.Number,
			Customer:   o.UserFullName,
			StatusCode: orderStatusCode,
			Total:      o.Total,
			CreatedAt:  timestamppb.New(o.CreatedAt),
			Product:    products,
		})
	}

	return &order.ListOrderAdminResponse{
		Base:       utils.SuccessResponse("get list order admin success"),
		Pagination: metadata,
		Items:      items,
	}, nil
}

func (os *orderService) ListOrder(ctx context.Context, request *order.ListOrderRequest) (*order.ListOrderResponse, error) {
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	orders, metadata, err := os.orderRepository.GetListOrderPagination(ctx, request.Pagination, claims.Subject)
	if err != nil {
		return nil, err
	}

	items := make([]*order.ListOrderResponseItem, 0)
	for _, o := range orders {

		products := make([]*order.ListOrderResponseItemProduct, 0)
		for _, oi := range o.Items {
			products = append(products, &order.ListOrderResponseItemProduct{
				Id:       oi.ProductID,
				Name:     oi.ProductName,
				Price:    oi.ProductPrice,
				Quantity: oi.Quantity,
			})
		}

		orderStatusCode := o.OrderStatusCode
		if o.OrderStatusCode == entity.OrderStatusCodeUnpaid && time.Now().After(*o.ExpiredAt) {
			orderStatusCode = entity.OrderStatusCodeExpired
		}

		xenditUrl := ""
		if o.XenditInvoiceUrl != nil {
			xenditUrl = *o.XenditInvoiceUrl
		}

		items = append(items, &order.ListOrderResponseItem{
			Id:               o.ID,
			Number:           o.Number,
			Customer:         o.UserFullName,
			StatusCode:       orderStatusCode,
			Total:            o.Total,
			CreatedAt:        timestamppb.New(o.CreatedAt),
			Product:          products,
			XenditInvoiceUrl: xenditUrl,
		})
	}

	return &order.ListOrderResponse{
		Base:       utils.SuccessResponse("get list order success"),
		Pagination: metadata,
		Items:      items,
	}, nil
}

func (os *orderService) DetailOrder(ctx context.Context, request *order.DetailOrderRequest) (*order.DetailOrderResponse, error) {
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	orderEntity, err := os.orderRepository.GetOrderByID(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	if claims.Role != entity.UserRoleAdmin && claims.Subject != orderEntity.UserID {
		return &order.DetailOrderResponse{
			Base: utils.BadRequestResponse("user id is not match"),
		}, nil
	}

	notes := ""
	if orderEntity.Notes != nil {
		notes = *orderEntity.Notes
	}
	xenditInvoiceUrl := ""
	if orderEntity.XenditInvoiceUrl != nil {
		xenditInvoiceUrl = *orderEntity.XenditInvoiceUrl
	}

	orderStatusCode := orderEntity.OrderStatusCode
	if orderEntity.OrderStatusCode == entity.OrderStatusCodeUnpaid && time.Now().After(*orderEntity.ExpiredAt) {
		orderStatusCode = entity.OrderStatusCodeExpired
	}

	items := make([]*order.DetailOrderResponseItem, 0)
	for _, oi := range orderEntity.Items {
		items = append(items, &order.DetailOrderResponseItem{
			Id:       oi.ProductID,
			Name:     oi.ProductName,
			Price:    oi.ProductPrice,
			Quantity: oi.Quantity,
		})
	}

	return &order.DetailOrderResponse{
		Base:             utils.SuccessResponse("get order detail success"),
		Id:               orderEntity.ID,
		Number:           orderEntity.Number,
		UserFullName:     orderEntity.UserFullName,
		Address:          orderEntity.Address,
		PhoneNumber:      orderEntity.PhoneNumber,
		Notes:            notes,
		OrderStatusCode:  orderStatusCode,
		CreatedAt:        timestamppb.New(orderEntity.CreatedAt),
		XenditInvoiceUrl: xenditInvoiceUrl,
		Items:            items,
		Total:            orderEntity.Total,
		ExpiredAt:        timestamppb.New(*orderEntity.ExpiredAt),
	}, nil
}

func (os *orderService) UpdateOrderStatus(ctx context.Context, request *order.UpdateOrderStatusRequest) (*order.UpdateOrderStatusResponse, error) {
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	orderEntity, err := os.orderRepository.GetOrderByID(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	if orderEntity == nil {
		return &order.UpdateOrderStatusResponse{
			Base: utils.NotFoundResponse("order not found"),
		}, nil
	}

	if claims.Role != entity.UserRoleAdmin && orderEntity.UserID != claims.Subject {
		return &order.UpdateOrderStatusResponse{
			Base: utils.BadRequestResponse("user id not match"),
		}, nil
	}

	if request.NewStatusCode == entity.OrderStatusCodePaid {
		if claims.Role != entity.UserRoleAdmin || orderEntity.OrderStatusCode != entity.OrderStatusCodeUnpaid {
			return &order.UpdateOrderStatusResponse{
				Base: utils.BadRequestResponse("update status is not allowed"),
			}, nil
		}
	} else if request.NewStatusCode == entity.OrderStatusCanceled {
		if orderEntity.OrderStatusCode != entity.OrderStatusCodeUnpaid {
			return &order.UpdateOrderStatusResponse{
				Base: utils.BadRequestResponse("update status is not allowed"),
			}, nil
		}
	} else if request.NewStatusCode == entity.OrderStatusCodeShipped {
		if claims.Role != entity.UserRoleAdmin || orderEntity.OrderStatusCode != entity.OrderStatusCodePaid {
			return &order.UpdateOrderStatusResponse{
				Base: utils.BadRequestResponse("update status is not allowed"),
			}, nil
		}
	} else if request.NewStatusCode == entity.OrderStatusCodeDone {
		if orderEntity.OrderStatusCode != entity.OrderStatusCodeShipped {
			return &order.UpdateOrderStatusResponse{
				Base: utils.BadRequestResponse("update status is not allowed"),
			}, nil
		}
	} else {
		return &order.UpdateOrderStatusResponse{
			Base: utils.BadRequestResponse("invalid a new status code"),
		}, nil
	}

	now := time.Now()
	orderEntity.OrderStatusCode = request.NewStatusCode
	orderEntity.UpdatedAt = &now
	orderEntity.UpdatedBy = &claims.Subject

	err = os.orderRepository.UpdateOrder(ctx, orderEntity)
	if err != nil {
		return nil, err
	}

	return &order.UpdateOrderStatusResponse{
		Base: utils.SuccessResponse("update order is success"),
	}, nil
}

func NewOrderService(db *sql.DB, orderRepository repository.IOrderRepository, productRepository repository.IProductRepository) IOrderService {
	return &orderService{
		db:                db,
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}
