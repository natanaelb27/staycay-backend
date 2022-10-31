package main

import (
	"os"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/handlers"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/middlewares"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/routes"
	"github.com/gin-gonic/gin"
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
