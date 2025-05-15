package lecturer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func GetRoleOfThesis(c *gin.Context) {
	var rolethesis []entity.RoleOfThesis

	db := config.DB()
	db.Find(&rolethesis)
	c.JSON(http.StatusOK, &rolethesis)
}
