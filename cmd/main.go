package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/jonayrodriguez/usermanagement/internal/config"
	"github.com/jonayrodriguez/usermanagement/internal/log"
)

const (
	developmentEnv = "development"
	productionEnv  = "production"
)

func main() {

	config, err := config.Load("../config/dev.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//logger.LogFactory.GetLogger
	systemLog := initSystemLog(config.Logging)
	accessLog := initAccessLog(config.Logging)

	//accessLog := log.New(config.Logging.AccessLevel, config.Logging.AccessFile, config.Logging.AccessFormat, false)

	if strings.ToLower(config.Server.Environment) == productionEnv {
		gin.SetMode(gin.ReleaseMode)

	}
	router := gin.Default()

	// Add a ginzap middleware, which:
	// NOTE: This could have its own local package to be able to update it by adding appoptics
	router.Use(ginzap.Ginzap(&accessLog.Logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(&accessLog.Logger, true))

	v1 := router.Group("/api/v1")
	{
		//TODO - Move logic to controllers
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	server := fmt.Sprintf("%s%s%d", config.Server.Host, ":", config.Server.Port)
	systemLog.Logger.Info("Application is running on " + server)

	router.Run(server)
}

func initSystemLog(c config.Logging) *log.Logger {
	lb := new(log.LoggerBuilder).SetLoggerType(log.SystemLogType).SetFullFilePath(c.File).SetFormat(c.Format).SetLevel(c.Level).SetMaxSize(c.MaxSize).SetMaxBackup(c.MaxBackup).SetMaxAge(c.MaxAge).NeedCaller(true).Build()

	systemLogger, err := lb.GetLogger()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return systemLogger

}

func initAccessLog(c config.Logging) *log.Logger {
	lb := new(log.LoggerBuilder).SetLoggerType(log.AccessLogType).SetFullFilePath(c.AccessFile).SetFormat(c.AccessFormat).SetLevel(c.AccessLevel).SetMaxSize(c.MaxSize).SetMaxBackup(c.MaxBackup).SetMaxAge(c.MaxAge).NeedCaller(false).Build()
	accessLogger, err := lb.GetLogger()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return accessLogger

}
