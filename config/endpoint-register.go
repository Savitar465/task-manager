package config

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, init *Initialization) {

	issues := router.Group("/issues")
	{
		issues.POST("", init.IssueCtrl.Create)
		issues.GET("", init.IssueCtrl.GetAll)
		issues.DELETE("", init.IssueCtrl.Delete)
	}
}
