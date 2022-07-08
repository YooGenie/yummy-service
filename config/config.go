package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

const (
	ContextUserClaimKey = "userClaim"
	ContextDBKey      = "db"
	ContextLogDBKey   = "logdb"
	ContextLogKey     = "log"
)

var Config = struct {
	HttpPort    string
	Environment string
	Jwt         struct {
		JwtSecret  string
		ContextKey string
	}
	Encrypt struct {
		EncryptKey string
	}
	Database struct {
		Driver           string
		User             string
		Connection       string
		ConnectionString string
		ShowSql          bool
	}
	Domain struct {
		WebApp struct {
			LoginUrl string
			HomeUrl  string
		}
		Service struct {
			ShareUrl string
		}
	}
	Log struct {
		ShowSql     bool
		ShowHttpLog bool
		Path        string
		MaxSize     int
		MaxBackups  int
		MaxAge      int
		Compress    bool
	}
}{}

func ConfigureEnvironment(path string, env ...string) {
	configor.Load(&Config, path+"/config.json")

	properties := make(map[string]string)

	for _, key := range env {
		arg := os.Getenv(key)
		if len(arg) == 0 {
			panic(fmt.Errorf("No %s system env variable\n", key))
		}
		properties[key] = arg
	}

	Config.Jwt.JwtSecret = properties["JWT_SECRET"]
	if Config.Jwt.JwtSecret == "" {
		Config.Jwt.JwtSecret = "TEST"
	}

	if properties["DAILY_WORK_LOG_DB_PASSWORD"] != "" {
		Config.Database.ConnectionString = fmt.Sprintf("%s:%s%s", Config.Database.User, properties["DAILY_WORK_LOG_DB_PASSWORD"], Config.Database.Connection)
	} else {
		Config.Database.ConnectionString = Config.Database.Connection
	}


	Config.Encrypt.EncryptKey = properties["DAILY_WORK_LOG_ENCRYPT_KEY"]

}