package databases

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/yescorihuela/deuna-payment-system/internal/shared/utils"

	_ "github.com/lib/pq"
)

func NewPostgresqlDbConnection(config utils.Config) (*sql.DB, error) {
	psql, err := sql.Open("postgres", config.DeunaDbDsn)
	if err != nil {
		panic(err)
	}
	migrator, err := postgres.WithInstance(psql, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", config.PathToMigrations), "deuna_payments", migrator)
	if err != nil {
		panic(err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
	return psql, nil
}
