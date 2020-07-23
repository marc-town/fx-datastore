package router

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	log.Print("Init router")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// ping
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is Running!")
	})

	// Global prefix
	g := e.Group("api/v1")

	// Router
	TransactionRouter(g)

	// Start erver
	e.Logger.Fatal(e.Start(":1323"))
}
