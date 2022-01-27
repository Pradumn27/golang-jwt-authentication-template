package routes

import (
	controller "github.com/Pradumn27/golang-jwt-authentication-template/controllers"

	"github.com/Pradumn27/golang-jwt-authentication-template/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/user", controller.GetUsers())
	incomingRoutes.GET("/user/:user_id", controller.GetUser())
}
