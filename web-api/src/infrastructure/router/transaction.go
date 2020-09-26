package router

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/marc-town/fx-datastore/web-api/src/infrastructure"
	"github.com/marc-town/fx-datastore/web-api/src/interfaces/controllers"
)

func TransactionRouter(g *echo.Group) {
	log.Print("Transaction router")

	// Controllers
	transactionController := controllers.NewTransactionController(infrastructure.NewSqlHandler())

	g.GET("/transactions", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "transactions")
	})
	// routting endpoint
	g.GET("/transactions", func(c echo.Context) error { return transactionController.Index(c) })
	g.GET("/transactions/:id", func(c echo.Context) error { return transactionController.Show(c) })
	g.POST("/transactions", func(c echo.Context) error { return transactionController.Create(c) })
	g.PUT("/transactions/:id", func(c echo.Context) error { return transactionController.Save(c) })
	g.DELETE("/transactions/:id", func(c echo.Context) error { return transactionController.Delete(c) })
}
