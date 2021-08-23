package main

import (
	"github.com/gin-gonic/gin"
	"github.com/s1ntaxe770r/opta-sanbox/db"
	_ "github.com/s1ntaxe770r/opta-sanbox/docs"
	"github.com/s1ntaxe770r/opta-sanbox/handlers"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Ekow API
// @version 1.0
// @description Ekow API documentation
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host http://prod.evil.corp
// @BasePath /
func main() {
	db.Connect()
	r := gin.Default()
	r.GET("/healthz", handlers.HealthCheck)
	r.POST("/set", handlers.Insert)
	r.POST("/get", handlers.Get)
	url := ginSwagger.URL("http://prod.evil.corp/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	err := r.Run(":8080")
	if err != nil {
		logrus.Fatal(err)
	}
}
