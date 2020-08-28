package Routes

import (
	"github.com/bobyhw39/rest_mvc/Controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine  {
	r:= gin.Default()
	grp1 := r.Group("/api/v1")
	{
		grp1.GET("account", Controller.GetAccount)
	}
	return r
}