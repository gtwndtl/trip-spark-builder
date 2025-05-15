package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetStatusRequest(c *gin.Context) {
	var StatusRequest []entity.StatusRequest

	db := config.DB()
	db.Find(&StatusRequest)
	c.JSON(http.StatusOK, &StatusRequest)
}