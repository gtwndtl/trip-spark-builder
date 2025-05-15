package condition

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func ListCondition(c *gin.Context) {
	var condition []entity.Condition

	db := config.DB()

	db.Find(&condition)

	c.JSON(http.StatusOK, &condition)
}