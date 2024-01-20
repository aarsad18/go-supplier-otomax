package model

type Transaction struct {
	ID                  string `gorm:"primaryKey"`
	TrxID               string `gorm:"column:trx_id"`
	SupplierTrxID       string `gorm:"column:supplier_trx_id"`
	AgensID             string `gorm:"column:id_agens"`
	CustID              string `gorm:"column:no_pelanggan"`
	ProductID           string `gorm:"column:id_produk"`
	ProductCode         string `gorm:"column:kode_produk"`
	SupplierID          string `gorm:"column:id_supplier"`
	Price               uint   `gorm:"column:harga_jual"`
	ReffSN              string `gorm:"column:reff_sn"`
	SupplierProductCode string `gorm:"column:kode_produk_supplier"`
	OrderID             string `gorm:"column:order_id"`
	LogIP               string `gorm:"column:log_ip"`
}

func (Transaction) TableName() string {
	return "t_transaksi"
}
