package model

type Transaction struct {
	ID                  string `gorm:"primaryKey"`
	TrxID               string `gorm:"column:trx_id"`
	SupplierTrxID       string `gorm:"column:supplier_trx_id"`
	CustID              string `gorm:"column:no_pelanggan"`
	ProductID           string `gorm:"column:id_produk"`
	ProductCode         string `gorm:"column:kode_produk"`
	SupplierID          string `gorm:"column:supplier_id"`
	ReffSN              string `gorm:"column:reff_sn"`
	SupplierProductCode string `gorm:"column:kode_produk_supplier"`
}

func (Transaction) TableName() string {
	return "t_transaksi"
}
