package repo

import (
	"log"

	"github.com/aarsad18/go-supplier-otomax/model"
	"github.com/aarsad18/go-supplier-otomax/resource"
)

type ISupplierRepo interface {
	FindByPK(id string) (*model.Supplier, error)
}

type SupplierRepo struct {
	db *resource.DBConn
}

func NewSupplierRepo(db *resource.DBConn) *SupplierRepo {
	return &SupplierRepo{db}
}

func (r *SupplierRepo) FindByPK(id string) (*model.Supplier, error) {

	var supplier model.Supplier
	err := r.db.DBGorm.Where("id = ?", id).Limit(1).Find(&supplier).Error
	if err != nil {
		log.Fatal(err)
	}

	return &supplier, err
}
