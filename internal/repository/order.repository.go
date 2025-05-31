package repository

import (
	"context"
	"database/sql"
	"errors"
	"github/eggnocent/app-grpc-eccomerce/internal/entity"
	"github/eggnocent/app-grpc-eccomerce/pkg/database"
	"log"
)

type IOrderRepository interface {
	WithTransaction(tx *sql.Tx) IOrderRepository
	GetNumbering(ctx context.Context, module string) (*entity.Numbering, error)
	CreateOrder(ctx context.Context, order *entity.Order) error
	UpdateNumbering(ctx context.Context, numbering *entity.Numbering) error
	CreateOrderItem(ctx context.Context, orderItem *entity.OrderItem) error
	GetOrderByID(ctx context.Context, orderID string) (*entity.Order, error)
	UpdateOrder(ctx context.Context, order *entity.Order) error
}

type orderRepository struct {
	db database.DatabaseQuery
}

func (os *orderRepository) WithTransaction(tx *sql.Tx) IOrderRepository {
	return &orderRepository{
		db: tx,
	}
}

func (or *orderRepository) GetNumbering(ctx context.Context, module string) (*entity.Numbering, error) {
	query := `
		SELECT 
			module, 
			number 
		FROM 
			numbering 
		WHERE 
			module = $1
		FOR
			UPDATE
	`

	row := or.db.QueryRowContext(ctx, query, module)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var numbering entity.Numbering
	err := row.Scan(
		&numbering.Module,
		&numbering.Numbering,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &numbering, nil

}

func (or *orderRepository) CreateOrder(ctx context.Context, order *entity.Order) error {
	query := `
		INSERT INTO "order" (
			id, 
			number, 
			user_id, 
			order_status_code, 
			user_full_name, 
			address, 
			phone_number, 
			notes, 
			total, 
			expired_at, 
			created_at, 
			created_by, 
			updated_at, 
			updated_by, 
			deleted_at, 
			deleted_by, 
			is_deleted,
			xendit_invoice_id,
			xendit_invoice_url
		) VALUES (
		 	$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13,
			$14,
			$15,
			$16,
			$17,
			$18,
			$19
		)
	`

	_, err := or.db.ExecContext(ctx, query,
		order.ID,
		order.Number,
		order.UserID,
		order.OrderStatusCode,
		order.UserFullName,
		order.Address,
		order.PhoneNumber,
		order.Notes,
		order.Total,
		order.ExpiredAt,
		order.CreatedAt,
		order.CreatedBy,
		order.UpdatedAt,
		order.UpdatedBy,
		order.DeletedAt,
		order.DeletedBy,
		order.IsDeleted,
		order.XenditInvoiceID,
		order.XenditInvoiceUrl,
	)

	if err != nil {
		return err
	}

	return nil
}

func (or *orderRepository) CreateOrderItem(ctx context.Context, orderItem *entity.OrderItem) error {
	query := `
		INSERT INTO order_item(
			id, 
			product_id,
			product_name, 
			product_image_file_name,
			product_price,
			quantity,
			order_id,
			created_at,
			created_by,
			updated_at,
			updated_by,
			deleted_at,
			deleted_by,
			is_deleted
		) VALUES (
			 $1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13,
			$14
		)
	`

	_, err := or.db.ExecContext(ctx, query,
		orderItem.ID,
		orderItem.ProductID,
		orderItem.ProductName,
		orderItem.ProductImageFileName,
		orderItem.ProductPrice,
		orderItem.Quantity,
		orderItem.OrderID,
		orderItem.CreatedAt,
		orderItem.CreatedBy,
		orderItem.UpdatedAt,
		orderItem.UpdatedBy,
		orderItem.DeletedAt,
		orderItem.DeletedBy,
		orderItem.IsDeleted,
	)

	if err != nil {
		return err
	}

	return nil
}

func (or *orderRepository) UpdateNumbering(ctx context.Context, numbering *entity.Numbering) error {
	query := `
		UPDATE 
			numbering 
		SET 
			number = $1 
		WHERE 
			module = $2
	`

	_, err := or.db.ExecContext(ctx, query,
		numbering.Numbering,
		numbering.Module,
	)

	if err != nil {
		return err
	}

	return nil
}

func (or *orderRepository) GetOrderByID(ctx context.Context, orderID string) (*entity.Order, error) {
	query := `
		SELECT 
			id
		FROM
			"order"
		WHERE
			id = $1
		AND
			is_deleted = false
	`

	row := or.db.QueryRowContext(ctx, query, orderID)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var order entity.Order
	err := row.Scan(
		&order.ID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}

func (or *orderRepository) UpdateOrder(ctx context.Context, order *entity.Order) error {
	log.Println("[REPOSITORY] Menjalankan UpdateOrder untuk ID:", order.ID)

	query := `
		UPDATE 
			"order"
		SET
			updated_at = $1,
			updated_by = $2,
			xendit_paid_at = $3,
			xendit_payment_channel = $4,
			xendit_payment_method = $5,
			order_status_code = $6
		WHERE	
			id = $7
	`

	_, err := or.db.ExecContext(ctx, query,
		order.UpdatedAt,
		order.UpdatedBy,
		order.XenditPaidAt,
		order.XenditPaymentChannel,
		order.XenditPaymentMethod,
		order.OrderStatusCode,
		order.ID, // Pastikan ini ditambahkan!
	)

	if err != nil {
		log.Println("[REPOSITORY] Gagal update order:", err)
		return err
	}

	log.Println("[REPOSITORY] Update order berhasil:", order.ID)
	return nil
}

func NewOrderRepository(db database.DatabaseQuery) IOrderRepository {
	return &orderRepository{
		db: db,
	}
}
