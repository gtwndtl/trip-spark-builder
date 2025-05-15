package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetSemester(c *gin.Context) {
	var semester []entity.Semester

	db := config.DB()
	db.Find(&semester)
	c.JSON(http.StatusOK, &semester)
}
