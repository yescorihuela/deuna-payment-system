package repositories

import (
	"database/sql"
	"fmt"

	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/transaction"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/models"
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

func (r *PostgresqlTransactionRepository) Create(transaction entities.Transaction) (*models.Transaction, error) {
	transactionModel := models.NewTransaction()
	query := shared.Compact(`
		INSERT INTO transactions
			(id, merchant_id, amount, status, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, merchant_id, amount, status, created_at`)
	err := r.db.QueryRow(query,
		transaction.Id,
		transaction.MerchantCode,
		transaction.Amount,
		transaction.Status,
		transaction.CreatedAt,
	).Scan(&transactionModel.Id, &transactionModel.MerchantId, &transactionModel.Amount, &transactionModel.Status, &transactionModel.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &transactionModel, nil
}

func (r *PostgresqlTransactionRepository) SetTransactionStatus(merchantCode, transactionId, status string) error {
	fmt.Println(merchantCode, transactionId, status)
	query := shared.Compact(`
		UPDATE transactions
		SET status = $1
		WHERE merchant_id = $2 AND id = $3`)
	_, err := r.db.Exec(query, status, merchantCode, transactionId)
	fmt.Println(merchantCode, transactionId)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresqlTransactionRepository) GetPaymentByTransactionId(merchantCode, transactionId string) (*models.Transaction, error) {
	transactionModel := models.NewTransaction()
	query := shared.Compact(`
					SELECT 
						id, merchant_id, amount, status, created_at
					FROM transactions
						WHERE merchant_id = $1 and id = $2`)

	err := r.db.QueryRow(query, merchantCode, transactionId).Scan(
		&transactionModel.Id,
		&transactionModel.MerchantId,
		&transactionModel.Amount,
		&transactionModel.Status,
		&transactionModel.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &transactionModel, nil
}

func (r *PostgresqlTransactionRepository) GetAllTransactionsByMerchant(merchantCode string) ([]*models.Transaction, error) {
	transactions := make([]*models.Transaction, 0)
	query := shared.Compact(`
		SELECT 
			id, merchant_id, amount, status, created_at
		FROM transactions
		WHERE merchant_id = $1
	`)
	results, err := r.db.Query(query, merchantCode)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		var transaction models.Transaction
		if err := results.Scan(&transaction.Id, &transaction.MerchantId, &transaction.Amount, &transaction.Status, &transaction.CreatedAt); err != nil {
			return transactions, err
		}
		transactions = append(transactions, &transaction)
	}

	return transactions, nil

}
