package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kuldeep/golang-jwt-project/utils"
)

func AuthRoutes(incomingRouts *gin.Engine) {
	incomingRouts.POST("users/signup", utils.Signup())
	incomingRouts.POST("users/login", utils.Login())
}
