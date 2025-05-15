package activity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetStatusActivity(c *gin.Context) {
	var statusActivity []entity.StatusActivity

	db := config.DB()
	db.Find(&statusActivity)
	c.JSON(http.StatusOK, &statusActivity)
}