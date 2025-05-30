package repository

import (
	"context"
	"database/sql"
	"errors"
	"github/eggnocent/app-grpc-eccomerce/internal/entity"
	"github/eggnocent/app-grpc-eccomerce/pkg/database"
)

type IOrderRepository interface {
	WithTransaction(tx *sql.Tx) IOrderRepository
	GetNumbering(ctx context.Context, module string) (*entity.Numbering, error)
	CreateOrder(ctx context.Context, order *entity.Order) error
	UpdateNumbering(ctx context.Context, numbering *entity.Numbering) error
	CreateOrderItem(ctx context.Context, orderItem *entity.OrderItem) error
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
			$14,
			$15,
			$16,
			$17
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

func NewOrderRepository(db database.DatabaseQuery) IOrderRepository {
	return &orderRepository{
		db: db,
	}
}
