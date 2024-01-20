package model

type DK string

const (
	KREDIT DK = "K"
	DEBIT  DK = "D"
)

type Deposit struct {
	ID           uint   `gorm:"primaryKey"`
	UserUpdate   string `gorm:"column:user_update"`
	AgensID      string `gorm:"column:id_agens"`
	TrxID        string `gorm:"column:trx_id"`
	SaldoAwal    uint   `gorm:"column:saldo_awal"`
	SaldoAkhir   uint   `gorm:"column:saldo_akhir"`
	Nominal      uint   `gorm:"column:nominal"`
	DK           DK     `gorm:"column:dk"`
	LastLog      string `gorm:"column:last_log"`
	MutationType string `gorm:"column:jenis_mutasi"`
}

func (Deposit) TableName() string {
	return "t_deposit"
}
