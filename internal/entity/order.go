package entity

import "time"

const (
	OrderStatusCodeUnpaid  = "unpaid"
	OrderStatusCodePaid    = "paid"
	OrderStatusCodeShipped = "shipped"
	OrderStatusCodeDone    = "done"
	OrderStatusCodeExpired = "expired"
	OrderStatusCanceled    = "canceled"
)

type Order struct {
	ID                   string
	Number               string
	UserID               string
	OrderStatusCode      string
	UserFullName         string
	Address              string
	PhoneNumber          string
	Notes                *string
	Total                float64
	ExpiredAt            *time.Time
	CreatedAt            time.Time
	CreatedBy            string
	UpdatedAt            *time.Time
	UpdatedBy            *string
	DeletedAt            *time.Time
	DeletedBy            *string
	IsDeleted            bool
	XenditInvoiceID      *string
	XenditInvoiceUrl     *string
	XenditPaidAt         *time.Time
	XenditPaymentMethod  *string
	XenditPaymentChannel *string

	Items []*OrderItem
}

type OrderItem struct {
	ID                   string
	ProductID            string
	ProductName          string
	ProductImageFileName string
	ProductPrice         float64
	Quantity             int64
	OrderID              string
	CreatedAt            time.Time
	CreatedBy            string
	UpdatedAt            *time.Time
	UpdatedBy            *string
	DeletedAt            *time.Time
	DeletedBy            *string
	IsDeleted            bool
}
