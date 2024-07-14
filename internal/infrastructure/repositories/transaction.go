package repositories

import (
	"database/sql"

	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/transaction"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/shared"
)

type PostgresqlTransactionRepository struct {
	db *sql.DB
}

func NewPostgresqlTransactionRepository(db *sql.DB) transaction.TransactionRepository {
	return &PostgresqlTransactionRepository{
		db: db,
	}
}

func (r *PostgresqlTransactionRepository) Create(refund entities.Transaction) (*entities.Transaction, error) {
	return &entities.Transaction{}, nil
}

func (r *PostgresqlTransactionRepository) SetTransactionStatus(merchantCode, transactionId, status string) error {
	query := shared.Compact(`
		UPDATE transactions
		SET status = $1
		WHERE merchantCode = $2 AND id = $3
	`)
	err := r.db.QueryRow(query, status, merchantCode, transactionId).Scan()
	if err != nil {
		return err
	}
	return nil
}
