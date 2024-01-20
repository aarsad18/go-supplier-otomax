package repo

import (
	"log"

	"github.com/aarsad18/go-supplier-otomax/model"
	"github.com/aarsad18/go-supplier-otomax/resource"
)

type ITransactionRepo interface {
	FindByTrxID(id string) (*model.Transaction, error)
	UpdateStatusAndSN(trxID string, data model.SupplierResult) error
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

	return &transaction, err
}

func (r *TransactionRepo) UpdateStatusAndSN(trxID string, data model.SupplierResult) error {
	// Execute the statement
	dataUpdate := map[string]interface{}{"status": data.Status, "reff_sn": data.SN}
	err := r.db.DBGorm.Where("trx_id = ?", trxID).Updates(dataUpdate).Error
	if err != nil {
		log.Fatal(err)
	}

	return err
}
