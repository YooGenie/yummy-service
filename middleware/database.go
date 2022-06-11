package middleware

import (
	"fmt"
	"github.com/YooGenie/daily-work-log-service/config"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo/v4"
	"net/http"
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

func setXormSession(engine *xorm.Engine) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//db session
			session := engine.NewSession()
			defer session.Close()

			c.Set(config.ContextDBKey, session)

			switch c.Request().Method {
			case "GET":
				if err := next(c); err != nil {
					return err
				}
			case "POST", "PUT", "DELETE":
				if err := session.Begin(); err != nil {
					return err
				}
				if err := next(c); err != nil {
					session.Rollback()
					return err
				}
				if c.Response().Status >= 500 {
					session.Rollback()

					// 처리 결과 에러 발생 시 rollback만 처리하고 여기서는 에러를 반환하지 않음
					// 이미 http 결과(c.Resopnse)에 에러 관련 정보가 담겨 있음
					return nil
				}
				if err := session.Commit(); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
			}

			return nil
		}
	}
}