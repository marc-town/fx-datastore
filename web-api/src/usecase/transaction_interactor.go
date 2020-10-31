package usecase

import (
	"github.com/marc-town/fx-datastore/web-api/src/domain/model"
	"github.com/marc-town/fx-datastore/web-api/src/domain/repository"
)

type TransactionInteractor struct {
	TransactionRepository repository.TransactionRepository
}

func (interactor *TransactionInteractor) TransactionById(id int) (transaction model.Transaction, err error) {
	transaction, err = interactor.TransactionRepository.FindById(id)
	return
}

func (interactor *TransactionInteractor) Transactions() (transaction model.Transactions, err error) {
	transaction, err = interactor.TransactionRepository.FindAll()
	return
}

func (interactor *TransactionInteractor) Add(a model.Transaction) (err error) {
	err = interactor.TransactionRepository.Store(a)
	return
}

func (interactor *TransactionInteractor) Update(a model.Transaction) (err error) {
	err = interactor.TransactionRepository.Update(a)
	return
}

func (interactor *TransactionInteractor) DeleteById(a model.Transaction) (err error) {
	err = interactor.TransactionRepository.DeleteById(a)
	return
}
