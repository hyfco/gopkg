package bootstrap

import (
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	conf "github.com/hyfco/gopkg/conf"
)

func NewSQLClient(conf *conf.Bootstrap) (*sql.DB, func(), error) {
	sqlDB, err := sql.Open("mysql", conf.GetData().GetMySQL().GetDSN())
	if err != nil {
		log.Fatalf("open db connection  failed.err:[%v]", err)
	}

	return sqlDB, func() {
		if err := sqlDB.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
