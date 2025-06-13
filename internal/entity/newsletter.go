package entity

import "time"

type Newsletter struct {
	ID        string
	FullName  string
	Email     string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy *string
	DeletedAt *time.Time
	DeletedBy *string
	IsDeleted bool
}
