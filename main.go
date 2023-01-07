package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/natanaelb27/staycay-backend/db"
	"github.com/natanaelb27/staycay-backend/internal/handlers"
	"github.com/natanaelb27/staycay-backend/internal/middlewares"
	"github.com/natanaelb27/staycay-backend/internal/routes"
)

func main() {
	r := gin.Default()

	db.Connect()

	r.Static("/docs", "swaggerui")

	r.NoRoute(middlewares.NotFoundHandler)

	r.Use(middlewares.CORSMiddleware())

	h := handlers.InitAllHandlers()

	routes.InitAllRoutes(r, h)

	var err error

	if os.Getenv("ENV") == "deployment" {
		err = r.Run(":" + os.Getenv("PORT"))
	} else {
		err = r.Run(":8080")
	}

	if err != nil {
		panic(err)
	}
}
