package main

import (
	"fmt"
	"os"
	"strings"

	mapping "github.com/jonayrodriguez/usermanagement/internal/router"

	"github.com/gin-gonic/gin"
	"github.com/jonayrodriguez/usermanagement/internal/config"
	"github.com/jonayrodriguez/usermanagement/internal/database"
	"github.com/jonayrodriguez/usermanagement/internal/log"
)

const (
	productionEnv = "production"
)

// 1 and 2 are already reserved.
const (
	errorCodeConfig int = iota + 2
	errorCodelogger
	errorCodeDB
)

func main() {

	conf, err := config.Load("../../config/dev.yml")
	// if there is any error loading the configuration, then exit.
	if err != nil {
		fmt.Println(err)
		os.Exit(errorCodeConfig)
	}

	//Initialize 2 type of logs (system and access)
	systemLog := initSystemLog(conf.Logging)
	accessLog := initAccessLog(conf.Logging)

	// If it´s not production, then it will be treated as development env.
	if strings.ToLower(conf.Server.Environment) == productionEnv {
		gin.SetMode(gin.ReleaseMode)

	}
	// keep in mind when it should be used pointer vs copy
	err = database.InitDB(conf.Database)
	// if there is any error connecting/migrating to DB the configuration.
	if err != nil {
		fmt.Println(err)
		os.Exit(errorCodeDB)
	}

	router := mapping.InitRouter(accessLog)

	server := fmt.Sprintf("%s%s%d", conf.Server.Host, ":", conf.Server.Port)
	systemLog.Logger.Info("Application is running on " + server)
	router.Run(server)

	//TODO- Build a gracefully shutdown

}

func initSystemLog(c config.Logging) *log.Logger {
	lb := new(log.LoggerBuilder).SetLoggerType(log.SystemLogType).SetFullFilePath(c.File).SetFormat(c.Format).SetLevel(c.Level).SetMaxSize(c.MaxSize).SetMaxBackup(c.MaxBackup).SetMaxAge(c.MaxAge).NeedCaller(true).Build()
	systemLogger, err := lb.GetLogger()
	if err != nil {
		fmt.Println(err)
		os.Exit(errorCodelogger)

	}

	return systemLogger

}

func initAccessLog(c config.Logging) *log.Logger {
	lb := new(log.LoggerBuilder).SetLoggerType(log.AccessLogType).SetFullFilePath(c.AccessFile).SetFormat(c.AccessFormat).SetLevel(c.AccessLevel).SetMaxSize(c.MaxSize).SetMaxBackup(c.MaxBackup).SetMaxAge(c.MaxAge).NeedCaller(false).Build()
	accessLogger, err := lb.GetLogger()
	if err != nil {
		fmt.Println(err)
		os.Exit(errorCodelogger)
	}

	return accessLogger

}
