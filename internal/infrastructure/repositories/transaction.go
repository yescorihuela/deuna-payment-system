package repositories

import (
	"database/sql"

	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/transaction"
)

type PostgresqlTransactionRepository struct {
	db *sql.DB
}

func NewPostgresqlTransactionRepository(db *sql.DB) transaction.TransactionRepository {
	return &PostgresqlTransactionRepository{
		db: db,
	}
}
