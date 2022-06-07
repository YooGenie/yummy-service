package main

import (
	"fmt"
	"github.com/YooGenie/daily-work-log-service/config"
	"github.com/YooGenie/daily-work-log-service/controller"
	"github.com/YooGenie/daily-work-log-service/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
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

	e := middleware.ConfigureEcho()

	controller.MemberController{}.Init(e.Group("/api/members"))

	log.Info("업무일지 Service Server Started: Port=" + config.Config.HttpPort)
	e.Start(":" + config.Config.HttpPort)
}
