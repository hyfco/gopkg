package bootstrap

import (
	"database/sql"
	sqldriver "github.com/go-sql-driver/mysql"
	"github.com/hyfco/gopkg/conf"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewGORM(conf *conf.Bootstrap) *gorm.DB {
	c := conf.GetData().GetMySql()
	if c == nil {
		panic("mysql config required")
	}

	sqlConf, err := sqldriver.ParseDSN(c.GetDSN())
	if err != nil {
		panic(err)
	}

	if !c.ReadTimeout.IsValid() {
		c.ReadTimeout = durationpb.New(DefaultReadTimeout * time.Millisecond)
	}

	if !c.WriteTimeout.IsValid() {
		c.WriteTimeout = durationpb.New(DefaultWriteTimeout * time.Millisecond)
	}

	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = DefaultMaxIdleConns
	}

	if c.MaxConns == 0 {
		c.MaxConns = DefaultMaxOpenConns
	}

	if !c.MaxLifeTime.IsValid() {
		c.MaxLifeTime = durationpb.New(DefaultMaxLifeTime * time.Millisecond)
	}

	sqlConf.Params = map[string]string{"charset": Chatset}
	sqlConf.Net = "tcp"
	sqlConf.MultiStatements = true
	sqlConf.AllowNativePasswords = true
	sqlConf.ReadTimeout = c.ReadTimeout.AsDuration()
	sqlConf.WriteTimeout = c.WriteTimeout.AsDuration()
	sqlConf.Timeout = DefaultDialTimeout * time.Millisecond

	db, err := sql.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// configure driver
	db.SetConnMaxLifetime(c.MaxLifeTime.AsDuration())
	db.SetMaxIdleConns(int(c.MaxIdleConns))
	db.SetMaxOpenConns(int(c.MaxConns))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{ /*Logger: logger*/ })
	if err != nil {
		panic(err)
	}

	return gormDB
}
