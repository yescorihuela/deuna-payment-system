package repositories

import (
	"database/sql"
	"time"

	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/merchant"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/mappers"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/models"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/shared"
)

type PostgresqlMerchantRepository struct {
	db *sql.DB
}

func NewPostgresqlMerchantRepository(db *sql.DB) merchant.MerchantRepository {
	return &PostgresqlMerchantRepository{
		db: db,
	}
}

func (r *PostgresqlMerchantRepository) Create(merchant entities.Merchant) (*models.Merchant, error) {
	model := mappers.FromMerchantEntityToModel(merchant)
	query := shared.Compact(`
					INSERT INTO merchants
						(name, balance, notification_email, merchant_code, enabled, created_at, updated_at)
					VALUES
						($1, $2, $3, $4, $5, $6, $7)
					RETURNING *`)
	err := r.db.QueryRow(query,
		merchant.Name,
		merchant.Balance,
		merchant.NotificationEmail,
		merchant.MerchantCode,
		merchant.Enabled,
		merchant.CreatedAt,
		merchant.UpdatedAt,
	).Scan(&model.Id, &model.Name, &model.Balance, &model.NotificationEmail, &model.MerchantCode, &model.Enabled, &model.CreatedAt, &model.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &model, err
}

func (r *PostgresqlMerchantRepository) GetByMerchantCode(merchantCode string) (*models.Merchant, error) {
	merchantModel := models.NewMerchantModel()
	query := shared.Compact(`
					SELECT 
						id, name, balance, notification_email, merchant_code, enabled, created_at, updated_at
					FROM merchants
						WHERE merchant_code = $1
	`)

	err := r.db.QueryRow(query, merchantCode).Scan(
		&merchantModel.Id,
		&merchantModel.Name,
		&merchantModel.Balance,
		&merchantModel.NotificationEmail,
		&merchantModel.MerchantCode,
		&merchantModel.Enabled,
		&merchantModel.CreatedAt,
		&merchantModel.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &merchantModel, nil
}

func (r *PostgresqlMerchantRepository) GetById(id int) (*models.Merchant, error) {
	merchantModel := models.NewMerchantModel()
	query := shared.Compact(`
				SELECT 
					id, name, balance, notification_email, merchant_code, enabled, created_at, updated_at
				FROM merchants
					WHERE id = $1
	`)

	err := r.db.QueryRow(query, id).Scan(
		&merchantModel.Id,
		&merchantModel.Name,
		&merchantModel.Balance,
		&merchantModel.NotificationEmail,
		&merchantModel.MerchantCode,
		&merchantModel.Enabled,
		&merchantModel.CreatedAt,
		&merchantModel.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &merchantModel, nil
}

func (r *PostgresqlMerchantRepository) DisableMerchant(merchantCode string, isEnabled bool) error {
	merchantModel := models.NewMerchantModel()
	query := shared.Compact(`
					UPDATE merchants
						SET enabled = $1
					WHERE
						id = $2
					RETURNING *
				`)
	err := r.db.QueryRow(query, isEnabled, merchantCode).Scan(&merchantModel.Id, &merchantModel.Name, &merchantModel.Balance, &merchantModel.NotificationEmail, &merchantModel.MerchantCode, &merchantModel.Enabled, &merchantModel.CreatedAt, &merchantModel.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresqlMerchantRepository) Update(id int, merchant entities.Merchant) (*models.Merchant, error) {
	merchantModel := models.NewMerchantModel()
	query := shared.Compact(`
					UPDATE merchants
						SET 
						name = $1,
						balance = $2,
						notification_email = $3,
						merchant_code = $4,
						enabled = $5,
						created_at = $6,
						updated_at = $7
					WHERE
						id = $8
					RETURNING *
				`)
	err := r.db.QueryRow(query,
		merchant.Name,
		merchant.Balance,
		merchant.NotificationEmail,
		merchant.MerchantCode,
		merchant.Enabled,
		merchant.CreatedAt.UTC(),
		time.Now().UTC(),
		id,
	).Scan(&merchantModel.Id, &merchantModel.Name, &merchantModel.Balance, &merchantModel.NotificationEmail, &merchantModel.MerchantCode, &merchantModel.Enabled, &merchantModel.CreatedAt, &merchantModel.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &merchantModel, nil

}
