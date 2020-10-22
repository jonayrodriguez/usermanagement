package router

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/jonayrodriguez/usermanagement/internal/controller"
	"github.com/jonayrodriguez/usermanagement/internal/log"
)

func InitRouter(logger *log.Logger) *gin.Engine {
	router := gin.Default()

	// Add a ginzap middleware, which:
	// NOTE: This could have its own local package to be able to update it by adding appoptics
	router.Use(ginzap.Ginzap(&logger.Logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(&logger.Logger, true))

	//TODO - JWT token , authentication (roles...), CORS, ... (App name could be added as part of the rooting)
	v1 := router.Group("/api/v1")
	{
		//TODO- Implement missing CRUD operations
		v1.GET("/users", controller.FindUsers)
		v1.GET("/users/:username", controller.GetUser)
		v1.POST("/users", controller.CreateUser)
		v1.DELETE("/users/:username", controller.DeleteUser)
	}

	return router

}
