package common

import (
	"github.com/YooGenie/daily-work-log-service/config"
	"github.com/go-xorm/xorm"
	echo "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserClaim struct {
	RequestID string `json:"-" `
	ID        int64  `json:"id"`
	Roles     string `json:"role"`
	Name      string `json:"name"`
	Datetime  string `json:"datetime"`
	Token     string `json:"-" `
}

func getContextValue(c echo.Context, key string) (value interface{}, exist bool) {
	value = c.Get(key)

	exist = value != nil

	return
}

func GetDB(c echo.Context) *xorm.Session {
	value, exist := getContextValue(c, config.ContextDBKey)

	if !exist {
		panic("DB is not exist")
	}
	if db, ok := value.(*xorm.Session); ok {
		return db
	}

	panic("DB is not exist")
}

func GetUserClaim(c echo.Context) *UserClaim {
	value, exist := getContextValue(c, config.ContextUserClaimKey)

	if !exist {
		return nil
	}

	return value.(*UserClaim)
}

func Log(c echo.Context) *logrus.Entry {
	value, exist := getContextValue(c, config.ContextLogKey)
	if !exist {
		panic("Log is not exist")
	}

	return value.(*logrus.Entry)
}
