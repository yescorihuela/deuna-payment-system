package repositories

import (
	"database/sql"

	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
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

func (r *PostgresqlMerchantRepository) Create(merchant entities.Merchant) (*entities.Merchant, error) {
	return &entities.Merchant{}, nil
}

func (r *PostgresqlMerchantRepository) GetByMerchantCode(merchantCode string) (*entities.Merchant, error) {
	return &entities.Merchant{}, nil
}

func (r *PostgresqlMerchantRepository) GetById(id int) (*entities.Merchant, error) {
	return &entities.Merchant{}, nil
}

func (r *PostgresqlMerchantRepository) DisableMerchant(merchantCode string) error {
	return nil
}

func (r *PostgresqlMerchantRepository) Update(id int, merchant entities.Merchant) (*entities.Merchant, error) {
	return &entities.Merchant{}, nil
}
