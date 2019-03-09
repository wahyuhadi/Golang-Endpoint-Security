package main

import (
	"fmt"
	"net/http"

	"./db"
	"./models"
	"github.com/gin-gonic/gin"
)

// DBMigrate for Auto Migrate
func DBMigrate() { /* Auto Migrations */
	fmt.Println("[::] Migration Databases .....")
	dbms := db.GetDatabaseConnection() /* Get connction to database */
	defer dbms.Close()
	dbms.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Profile{}) /* Migration Models */
	fmt.Println("[::] Migration Databases Done")
}

func main() {
	DBMigrate()
	router := gin.Default()
	router.POST("/user", CreateUser)
	router.GET("/user", GetUser)
	router.Run(":8080")
}

// CreateUser function
func CreateUser(c *gin.Context) {
	var (
		user models.User
	)

	c.BindJSON(&user)
	dbms := db.GetDatabaseConnection() /*  Open connectins */
	defer dbms.Close()
	dbms.Create(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "result": user})
	return
}

// GetUser function
func GetUser(c *gin.Context) {
	var user []models.User
	dbms := db.GetDatabaseConnection() /*  Open connectins */
	defer dbms.Close()
	err := dbms.Find(&user).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "result": user})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "something error", "result": err})
	}
	return

}
