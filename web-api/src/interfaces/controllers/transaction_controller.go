package controllers

import (
	"log"
	"strconv"

	"github.com/labstack/echo"
	"github.com/marc-town/fx-datastore/web-api/src/domain/model"
	"github.com/marc-town/fx-datastore/web-api/src/interfaces/database"
	"github.com/marc-town/fx-datastore/web-api/src/usecase"
)

type TransactionController struct {
	Interactor usecase.TransactionInteractor
}

func NewTransactionController(sqlHandler database.SqlHandler) *TransactionController {
	return &TransactionController{
		Interactor: usecase.TransactionInteractor{
			TransactionRepository: &database.TransactionRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TransactionController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := controller.Interactor.TransactionById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, transaction)
	return
}

func (controller *TransactionController) Index(c echo.Context) (err error) {
	transactions, err := controller.Interactor.Transactions()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, transactions)
	return
}

func (controller *TransactionController) Create(c echo.Context) (err error) {
	transaction := model.Transaction{}
	c.Bind(&transaction)
	if err = controller.Interactor.Add(transaction); err != nil {
		log.Print(err)
		c.JSON(500, SetErrorResponse("error_code"))
		return
	}
	c.NoContent(201)
	return
}

func (controller *TransactionController) Save(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction := model.Transaction{
		ID: id,
	}
	c.Bind(&transaction)
	if err = controller.Interactor.Update(transaction); err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.NoContent(201)
	return
}

func (controller *TransactionController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction := model.Transaction{
		ID: id,
	}
	err = controller.Interactor.DeleteById(transaction)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.NoContent(200)
	return
}
