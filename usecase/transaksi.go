package usecase

import (
	"log"

	"github.com/aarsad18/go-supplier-otomax/model"
	"github.com/aarsad18/go-supplier-otomax/repo"
)

type ITransaksiUsecase interface {
	PulsaTrx(payload model.PgNotificationPayload) *model.Transaction
}

type RepoInit struct {
	SupplierRepo    repo.ISupplierRepo
	TransactionRepo repo.ITransactionRepo
	OtomaxRepo      repo.IOtomaxRepo
	DepositRepo     repo.IDepositRepo
}

type TransaksiUsecase struct {
	repo RepoInit
}

func NewTransaksiUsecase(repo RepoInit) *TransaksiUsecase {
	return &TransaksiUsecase{repo}
}

func (t *TransaksiUsecase) PulsaTrx(payload model.PgNotificationPayload) *model.Transaction {
	supplier, err := t.repo.SupplierRepo.FindByPK(payload.SupplierID)
	if err != nil {
		log.Print(err)
	}

	transaction, err := t.repo.TransactionRepo.FindByTrxID(payload.TrxID)
	if err != nil {
		log.Print(err)
	}

	oto, err := t.repo.OtomaxRepo.RequestTransaction(*supplier, *transaction)
	if err != nil {
		log.Print(err)
	}

	err = t.repo.TransactionRepo.UpdateStatusAndSN(transaction.TrxID, oto)
	if err != nil {
		log.Print(err)
	}

	if oto.Status == model.FAILED {
		log.Print("transaction failed")
		err := t.repo.DepositRepo.RefundSaldo(*transaction)
		if err != nil {
			log.Printf("Error refund saldo: %s", err)
		}
		log.Print("sukses refund saldo")
	} else {
		log.Printf("status transaksi %s", oto.Status)
	}

	return transaction
}
