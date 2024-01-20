package model

type PgNotificationPayload struct {
	ID              int64  `json:"id"`
	TrxID           string `json:"trx_id"`
	SupplierID      string `json:"id_supplier"`
	ProductCategory string `json:"produk_kategori"`
	IsTransaction   string `json:"is_transaksi"`
}
