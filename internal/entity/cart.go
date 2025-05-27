package entity

import "time"

type UserCart struct {
	ID        string
	UserID    string
	ProductID string
	Quantity  int
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy *string

	Product *Product
}
