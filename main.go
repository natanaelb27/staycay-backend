package main

import (
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/db"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/handlers"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/middlewares"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.Connect()

	r.Use(middlewares.CORSMiddleware())
	r.NoRoute(middlewares.NotFoundHandler)
	h := handlers.InitAllHandlers()

	routes.InitAllRoutes(r, h)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
