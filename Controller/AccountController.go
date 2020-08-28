package Controller

import (
	"github.com/bobyhw39/rest_mvc/Entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAccount(c *gin.Context)  {
	var account []Entity.Account
	err := Entity.GetAllAccounts(&account)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK,account)
	}
}