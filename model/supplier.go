package model

import "github.com/aarsad18/go-supplier-otomax/repo"

type Supplier struct {
	ID        string `gorm:"primaryKey"`
	Name      string `gorm:"column:nama"`
	Url       string
	Username  string
	PIN       string
	Signature string
	Key       string
}

func (Supplier) TableName() string {
	return "m_supplier"
}

type SupplierResult struct {
	Status repo.TrxStatus
	SN     string
	Msg    string
}
