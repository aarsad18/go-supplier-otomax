package repo

import (
	"log"

	"github.com/aarsad18/go-supplier-otomax/model"
	"github.com/aarsad18/go-supplier-otomax/resource"
)

type ITransactionRepo interface {
	FindByTrxID(id string) (*model.Transaction, error)
}

type TransactionRepo struct {
	db *resource.DBConn
}

func NewTransactionRepo(db *resource.DBConn) *TransactionRepo {
	return &TransactionRepo{db}
}

func (r *TransactionRepo) FindByTrxID(id string) (*model.Transaction, error) {
	// Execute the statement
	var transaction model.Transaction
	err := r.db.DBGorm.Where("trx_id = ?", id).Limit(1).Find(&transaction).Error
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("transaction data: %s\n", transaction)

	return &transaction, err
}
