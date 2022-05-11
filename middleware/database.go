package middleware

import (
	"fmt"
	"github.com/YooGenie/daily-work-log-service/config"
	"github.com/go-xorm/xorm"
	"time"
	"xorm.io/core"
)

var xormDb *xorm.Engine

func ConfigureDatabase() *xorm.Engine {
	var err error

	dbConnection := config.Config.Database.ConnectionString

	xormDb, err = xorm.NewEngine(config.Config.Database.Driver, dbConnection)
	if err != nil {
		panic(fmt.Errorf("Database open error: error: %s \n", err))
	}

	xormDb.SetMaxOpenConns(10)
	xormDb.SetMaxIdleConns(5)
	xormDb.SetConnMaxLifetime(10 * time.Minute)

	xormDb.ShowSQL(config.Config.Database.ShowSql)
	xormDb.Logger().SetLevel(core.LOG_INFO)

	return xormDb
}
