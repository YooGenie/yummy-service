package main

import (
	"fmt"
	"github.com/YooGenie/daily-work-log-service/config"
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

	fmt.Println("환경변수 설정 끝")
}