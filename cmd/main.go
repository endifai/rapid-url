package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"rapidurl/internal/config"
	"rapidurl/internal/controllers"
	"rapidurl/internal/database"
	"rapidurl/internal/repositories"
)

func main() {
	config.InitConfig()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, error := database.NewDatabase(config.REDIS_ADDRESS, config.REDIS_PASSWORD)

	if error != nil {
		panic(error)
	}

	linkRepository := repositories.NewLinkRepository(db)
	linkController := controllers.NewLinkController(linkRepository)

	e.POST("/api/shorten-url", linkController.CreateLink)
	e.GET("/s/:hash", linkController.GetLink)

	e.Logger.Fatal(e.Start(":" + config.PORT))
}
