package lecturer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetDegrees(c *gin.Context) {
	var degree []entity.Degree

	db := config.DB()
	db.Find(&degree)
	c.JSON(http.StatusOK, &degree)
}

func GetDegreesWithOutStudent(c *gin.Context) {
	var degree []entity.Degree

	db := config.DB()
	db.Where("id != 1").Find(&degree)
	c.JSON(http.StatusOK, &degree)
}
