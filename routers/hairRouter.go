package routers

import (
	"github.com/gin-gonic/gin"
	"mp_royal/controllers"
)

func SetHairRouters(router *gin.RouterGroup) {
	router.GET("/cycle", controllers.GetHairCycle)
	router.GET("/date", controllers.GetHairDate)
}
