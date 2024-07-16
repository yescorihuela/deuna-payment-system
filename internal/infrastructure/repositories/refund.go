package repositories

import (
	"database/sql"
	"errors"
	"fmt"

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
			(id, transaction_id, merchant_id, amount, status, created_at)
		VALUES
			($1, $2, $3, $4, $5, $6)
		RETURNING *;
	`)
	err := r.db.QueryRow(query,
		refund.Id,
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

func (r *PostgresqlRefundRepository) GetRefundByTransactionId(merchantCode, transactionId string) (bool, error) {
	var refundExists bool
	query := shared.Compact(`
	SELECT 
		EXISTS( 
			SELECT 
				1 
			FROM refunds 
			WHERE merchant_id = $1 AND transaction_id = $2			
		)`)
	err := r.db.QueryRow(query, merchantCode, transactionId).Scan(&refundExists)
	fmt.Println(err, refundExists)
	if errors.Is(err, sql.ErrNoRows) {
		return false, err
	}
	return refundExists, nil
}
