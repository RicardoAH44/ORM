package main

import (
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    uint    `gorm:"primaryKey" json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var db *gorm.DB

func main() {
	dsn := "RicardoAH:Tato9680@tcp(127.0.0.1:3306)/test?parseTime=true"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Product{})

	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProduct)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

	router.Run(":8080")
}

func getProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Product
	db.First(&product, id)
	c.JSON(http.StatusOK, product)
}

func createProduct(c *gin.Context) {
	var product Product
	c.BindJSON(&product)
	db.Create(&product)
	c.JSON(http.StatusCreated, product)
}

func updateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Product
	db.First(&product, id)
	if product.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&product)
	db.Save(&product)
	c.JSON(http.StatusOK, product)
}

func deleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Product
	db.Where("id = ?", id).Delete(&product)
	c.JSON(http.StatusNoContent, nil)
}
