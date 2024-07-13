package repositories

import (
	"database/sql"

	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/merchant"
)

type PostgresqlMerchantRepository struct {
	db *sql.DB
}

func NewPostgresqlMerchantRepository(db *sql.DB) merchant.MerchantRepository {
	return &PostgresqlMerchantRepository{
		db: db,
	}
}
