package payment


import (

   "net/http"


   "github.com/gin-gonic/gin"
	"github.com/sut67/team09/config"
	"github.com/sut67/team09/entity"


)


func GetAll(c *gin.Context) {


   db := config.DB()


   var payment []entity.Payment
   db.Preload("Users").Preload("Dormitory").Find(&payment)


   c.JSON(http.StatusOK, &payment)


}

// func Delete(c *gin.Context) {


//    id := c.Param("id")

//    db := config.DB()

//    if tx := db.Exec("DELETE FROM report WHERE id = ?", id); tx.RowsAffected == 0 {

//        c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})

//        return

//    }

//    c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

// }

func Get(c *gin.Context) {


   ID := c.Param("id")

   var payment entity.Payment


   db := config.DB()

   results := db.Preload("Users").Preload("Dormitory").First(&payment, ID)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   if payment.ID == 0 {

       c.JSON(http.StatusNoContent, gin.H{})

       return

   }

   c.JSON(http.StatusOK, payment)


}

func Update(c *gin.Context) {


    var payment entity.Payment
 
 
    PaymentID := c.Param("id")
 
 
    db := config.DB()
 
    result := db.First(&payment, PaymentID)
 
    if result.Error != nil {
 
        c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
 
        return
 
    }
 
 
    if err := c.ShouldBindJSON(&payment); err != nil {
 
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
 
        return
 
    }
 
 
    result = db.Save(&payment)
 
    if result.Error != nil {
 
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
 
        return
 
    }
 
 
    c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
 
 }


 func Create(c *gin.Context) {
    var newPayment entity.Payment

    // Bind the incoming JSON payload to the newAdmin struct
    if err := c.ShouldBindJSON(&newPayment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
        return
    }

    db := config.DB()

    // Save the new payment to the database
    if err := db.Create(&newPayment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Payment created successfully", "payment": newPayment})
}