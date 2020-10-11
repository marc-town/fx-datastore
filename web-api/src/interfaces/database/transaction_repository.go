package database

import "github.com/marc-town/fx-datastore/web-api/src/domain/model"

type TransactionRepository struct {
	SqlHandler
}

func (repo *TransactionRepository) FindById(id int) (transaction model.Transaction, err error) {
	if err = repo.Find(&transaction, id).Error; err != nil {
		return
	}
	return
}

func (repo *TransactionRepository) FindAll() (transactions model.Transactions, err error) {
	if err = repo.Find(&transactions).Error; err != nil {
		return
	}
	return
}

func (repo *TransactionRepository) Store(a model.Transaction) (err error) {
	if err = repo.Create(&a).Error; err != nil {
		return
	}
	return
}

func (repo *TransactionRepository) Update(a model.Transaction) (transaction model.Transaction, err error) {
	if err = repo.Save(&a).Error; err != nil {
		return
	}
	transaction = a
	return
}

func (repo *TransactionRepository) DeleteById(transaction model.Transaction) (err error) {
	if err = repo.Delete(&transaction).Error; err != nil {
		return
	}
	return
}
