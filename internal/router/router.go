package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jonayrodriguez/usermanagement/internal/controller"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

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
