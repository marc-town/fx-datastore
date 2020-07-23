package repository

import "github.com/marc-town/fx-datastore/web-api/src/domain/model"

type TransactionRepository interface {
	FindById(id int) (model.Transaction, error)
	FindAll() (model.Transactions, error)
	Store(model.Transaction) (model.Transaction, error)
	Update(model.Transaction) (model.Transaction, error)
	DeleteById(model.Transaction) error
}
