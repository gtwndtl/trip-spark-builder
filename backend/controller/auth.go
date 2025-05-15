package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"
	"github.com/sut67/team09/services"

	"golang.org/x/crypto/bcrypt"
)


type (

   Authen struct {

       Email    string `json:"email"`

       Password    string `json:"password"`

   }

)

func SignIn(c *gin.Context) {
    var payload Authen
    var lecturer entity.Lecturer

    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB().Raw("SELECT * FROM lecturers WHERE email = ?", payload.Email).Scan(&lecturer).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := bcrypt.CompareHashAndPassword([]byte(lecturer.Password), []byte(payload.Password))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "password is incorrect"})
        return
    }

    jwtWrapper := services.JwtWrapper{
        SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
        Issuer:          "AuthService",
        ExpirationHours: 24,
    }

    signedToken, err := jwtWrapper.GenerateToken(lecturer.Email)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
        return
    }

    
    c.JSON(http.StatusOK, gin.H{
        "token_type":  "Bearer",
        "token":       signedToken,
        "id":          lecturer.ID,
        "position_id": lecturer.PositionID,
        "Email": lecturer.Email, 
    })
}