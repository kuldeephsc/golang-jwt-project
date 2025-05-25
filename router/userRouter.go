package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kuldeep/golang-jwt-project/middleware"
	"github.com/kuldeep/golang-jwt-project/utils"
)

func UserRouts(incomingRouts *gin.Engine) {
	// Only protect user GET routes
	userGroup := incomingRouts.Group("/users")
	userGroup.Use(middleware.Authenticate())
	{
		userGroup.GET("", utils.GetUser())
		userGroup.GET("/:user_id", utils.GetUser())
	}
}
