package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jonayrodriguez/usermanagement/internal/config/config"
)

func main() {
	config.Load("", nil)
	// Logging to a file.
	f, _ := os.OpenFile("../logs/usermanagement.log", os.O_CREATE|os.O_APPEND, 0644)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
