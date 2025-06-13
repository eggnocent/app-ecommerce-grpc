package repository

import (
	"context"
	"database/sql"
	"errors"
	"github/eggnocent/app-grpc-eccomerce/internal/entity"
)

type INewsletterRepository interface {
	GetNewsletterByEmail(ctx context.Context, email string) (*entity.Newsletter, error)
	CreateNewNewsletter(ctx context.Context, newsletter *entity.Newsletter) error
}

type newsletterRepository struct {
	db sql.DB
}

func (nr *newsletterRepository) GetNewsletterByEmail(ctx context.Context, email string) (*entity.Newsletter, error) {
	query := `
		SELECT 
			id 
		FROM 
			newsletter 
		WHERE 
			email = $1
		AND 
			is_deleted = false
	`
	row := nr.db.QueryRowContext(ctx, query, email)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var newsletter entity.Newsletter
	err := row.Scan(
		&newsletter.ID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &newsletter, nil
}

func (nr *newsletterRepository) CreateNewNewsletter(ctx context.Context, newsletter *entity.Newsletter) error {
	query := `
		INSERT INTO newsletter (
			id,
			full_name,
			email,
			created_at,
			created_by
		) VALUES (
			$1, $2, $3, $4, $5 
		)
	`

	_, err := nr.db.ExecContext(ctx, query,
		newsletter.ID,
		newsletter.FullName,
		newsletter.Email,
		newsletter.CreatedAt,
		newsletter.CreatedBy,
	)

	if err != nil {
		return err
	}

	return nil
}

func NewNewsletterRepository(db *sql.DB) INewsletterRepository {
	return &newsletterRepository{
		db: *db,
	}
}
