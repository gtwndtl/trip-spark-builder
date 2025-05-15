package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetRequestTypes(c *gin.Context) {
	var RequestTypes []entity.RequestType

	db := config.DB()
	db.Find(&RequestTypes)
	c.JSON(http.StatusOK, &RequestTypes)
}