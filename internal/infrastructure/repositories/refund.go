package repositories

import (
	"database/sql"

	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/refund"
)

type PostgresqlRefundRepository struct {
	db *sql.DB
}

func NewPostgresqlRefundRepository(db *sql.DB) refund.RefundRepository {
	return &PostgresqlRefundRepository{
		db: db,
	}
}

func (r *PostgresqlRefundRepository) Create(refund entities.Refund) (*entities.Refund, error) {
	return &entities.Refund{}, nil
}
