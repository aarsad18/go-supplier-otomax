package model

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
	Status TrxStatus
	SN     string
	Msg    string
}
