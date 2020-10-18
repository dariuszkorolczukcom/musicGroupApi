package main

import (
	"github.com/dariuszkorolczukcom/musicGroupApi/pkg/handlers"
	DB "github.com/dariuszkorolczukcom/musicGroupApi/util/mongoDB"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	DB.InitDB()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/presets", handlers.GetPreset)
	e.GET("/presets/:id", handlers.GetPreset)
	e.POST("/presets", handlers.CreatePreset)
	e.PUT("/presets/:id", handlers.UpdatePreset)
	e.DELETE("/presets/:id", handlers.DeletePreset)
	e.Logger.Fatal(e.Start(":8080"))
}
