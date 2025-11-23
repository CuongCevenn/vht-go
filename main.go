package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"vht-go/middleware"
	categorymodule "vht-go/modules/category"
	restaurantmodule "vht-go/modules/restaurant"
	usermodule "vht-go/modules/user"
	sharedcomponent "vht-go/shared/component"

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

	expIn := 60 * 60 * 24 * 7

	if os.Getenv("JWT_EXP_IN") != "" {
		expInInt, err := strconv.Atoi(os.Getenv("JWT_EXP_IN"))
		if err != nil {
			log.Fatalln("failed to convert JWT_EXP_IN to int")
		}
		expIn = expInInt
	}

	jwtComponent := sharedcomponent.NewJwtComp(os.Getenv("JWT_SECRET_KEY"), expIn)

	tokenIntrospector := sharedcomponent.NewTokenIntrospector(jwtComponent, db)
	middlewareProvider := middleware.NewMiddlewareProvider(tokenIntrospector)

	r := gin.Default()

	r.Use(middleware.RecoverMiddleware())

	v1 := r.Group("/v1")

	categorymodule.SetupCategoryModule(v1, db)
	restaurantmodule.SetupRestaurantModule(v1, db)
	usermodule.SetupUserModule(v1, db, jwtComponent, middlewareProvider)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:3600
}
