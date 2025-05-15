package lecturer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetGenders(c *gin.Context) {
	var genders []entity.Gender

	db := config.DB()
	db.Find(&genders)
	c.JSON(http.StatusOK, &genders)
}
