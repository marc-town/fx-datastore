package model

type Transactions []Transaction

type Transaction struct {
	ID    int
	Name  string
	Price int
}
