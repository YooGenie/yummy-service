package main

import (
	"fmt"
	"github.com/YooGenie/daily-work-log-service/config"
	"github.com/YooGenie/daily-work-log-service/middleware"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	config.ConfigureEnvironment("./config", "JWT_SECRET", "DAILY_WORK_LOG_DB_PASSWORD", "DAILY_WORK_LOG_ENCRYPT_KEY")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()

	xorm := middleware.ConfigureDatabase()
	xorm.Close()
	fmt.Println(xorm)
	//e := echo.New()

	fmt.Println("DB 설정 끝")
}
