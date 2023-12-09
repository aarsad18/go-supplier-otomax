package usecase

import (
	"fmt"
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
		log.Fatal(err)
	}

	transaction, err := t.repo.TransactionRepo.FindByTrxID(payload.TrxID)
	if err != nil {
		log.Fatal(err)
	}

	oto, err := t.repo.OtomaxRepo.RequestTransaction(*supplier, *transaction)
	if err != nil {
		log.Fatal(err)
	}

	if oto.Status == repo.SUCCESS {
		// update SN
	} else if oto.Status == repo.FAILED {
		// refund saldo
	}

	fmt.Println(oto)

	return transaction
}
