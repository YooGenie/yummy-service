package context

import (
	"github.com/YooGenie/daily-work-log-service/config"
	"github.com/go-xorm/xorm"
	echo "github.com/labstack/echo/v4"
)


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


