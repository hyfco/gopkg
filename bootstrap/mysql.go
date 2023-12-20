package bootstrap

import (
	"database/sql"
	entSql "entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	conf "github.com/hyfco/gopkg/config"
)

func NewEntDriver(conf *conf.Data, logger log.Logger) (*entSql.Driver, func(), error) {
	sqlDB, err := sql.Open(conf.Database.Driver, conf.Database.Source)
	if err != nil {
		log.Fatalf("open db connection  failed.err:[%v]", err)
	}

	driver := entSql.OpenDB(conf.Database.Driver, sqlDB)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	return driver, func() {
		if err := sqlDB.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func NewSQLClient(conf *conf.Data, logger log.Logger) (*sql.DB, func(), error) {
	sqlDB, err := sql.Open(conf.Database.Driver, conf.Database.Source)
	if err != nil {
		log.Fatalf("open db connection  failed.err:[%v]", err)
	}

	return sqlDB, func() {
		if err := sqlDB.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
