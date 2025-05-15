package lecturer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetPositions(c *gin.Context) {
	var position []entity.Position

	db := config.DB()
	db.Find(&position)
	c.JSON(http.StatusOK, &position)
}

func GetPositionLecturer(c *gin.Context) {
	var position []entity.Position

	db := config.DB()
	db.Where("id != ?", 7).Find(&position)
	c.JSON(http.StatusOK, &position)
}
