package repo

import (
	"fmt"
	"log"

	"github.com/aarsad18/go-supplier-otomax/model"
	"github.com/aarsad18/go-supplier-otomax/resource"
)

type IDepositRepo interface {
	RefundSaldo(data model.Transaction) error
}

type DepositRepo struct {
	db *resource.DBConn
}

func NewDepositRepo(db *resource.DBConn) *DepositRepo {
	return &DepositRepo{db}
}

func (r *DepositRepo) RefundSaldo(data model.Transaction) error {

	lastLog := fmt.Sprintf("%s | %s", data.OrderID, data.LogIP)

	var params []interface{}

	params = append(params, data.Price, data.Price, lastLog, data.TrxID, data.AgensID)

	err := r.db.DBGorm.Exec(`
		UPDATE public.t_deposit
		SET 
			time_update=CURRENT_TIMESTAMP, 
			user_update=?,
			saldo_awal=saldo_akhir, 
			saldo_akhir=saldo_akhir + ?, 
			nominal=?, 
			dk='K', 
			last_log=?, 
			trx_id=?, 
			jenis_mutasi='CTR'
		WHERE
			id_agens = ?`, "APP-SUPPLIER", params).Error
	if err != nil {
		log.Fatal(err)
	}

	return err
}
