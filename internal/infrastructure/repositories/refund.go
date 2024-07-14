package repositories

import (
	"database/sql"

	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/refund"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/models"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/shared"
)

type PostgresqlRefundRepository struct {
	db *sql.DB
}

func NewPostgresqlRefundRepository(db *sql.DB) refund.RefundRepository {
	return &PostgresqlRefundRepository{
		db: db,
	}
}

func (r *PostgresqlRefundRepository) Create(refund entities.Refund) (*models.Refund, error) {
	refundModel := models.NewRefund()
	query := shared.Compact(`
		INSERT INTO refunds
			(transaction_id, merchant_id, amount, status, created_at)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING *;
	`)
	err := r.db.QueryRow(query,
		refund.TransactionId,
		refund.MerchantId,
		refund.Amount,
		refund.Status,
		refund.CreatedAt,
	).Scan(
		&refundModel.Id,
		&refundModel.TransactionId,
		&refundModel.MerchantId,
		&refundModel.Amount,
		&refundModel.Status,
		&refundModel.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &refundModel, nil
}
