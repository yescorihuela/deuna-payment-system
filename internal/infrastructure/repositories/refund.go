package repositories

import (
	"database/sql"

	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/refund"
)

type PostgresqlRefundRepository struct {
	db *sql.DB
}

func NewPostgresqlRepository(db *sql.DB) refund.RefundRepository {
	return &PostgresqlRefundRepository{
		db: db,
	}
}
