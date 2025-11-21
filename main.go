package main

import (
	"fmt"
	"log"
	"os"
	categorymodule "vht-go/modules/category"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)	

func main() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("failed to connect database")
	}
	fmt.Println("Database connected")

	r := gin.Default()

	v1 := r.Group("/v1")

	catGroup := v1.Group("/categories")
	{
		catGroup.POST("", categorymodule.CreateCategory(db))
		catGroup.GET("/:id", categorymodule.GetCategoryById(db))
		catGroup.GET("", categorymodule.ListCategories(db))
		catGroup.PATCH("/:id", categorymodule.UpdateCategory(db))
		catGroup.DELETE("/:id", categorymodule.DeleteCategory(db))
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:3600
}