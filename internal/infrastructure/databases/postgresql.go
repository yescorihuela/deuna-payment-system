package databases

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

func NewPostgresqlDbConnection() (*sql.DB, error) {
	psql, err := sql.Open("postgres", "") // Use envvars or viper
	if err != nil {
		panic(err)
	}
	migrator, err := postgres.WithInstance(psql, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	migration, err := migrate.NewWithDatabaseInstance("", "postgres", migrator)
	if err != nil {
		panic(err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
	return psql, nil
}
