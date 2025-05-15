package studentevaluation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
)

func ListAssessmentType(c *gin.Context) {
	var assessmenttype []entity.AssessmentType

	db := config.DB()

	db.Find(&assessmenttype)

	c.JSON(http.StatusOK, &assessmenttype)
}