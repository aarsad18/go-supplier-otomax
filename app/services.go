package app

import (
	"github.com/aarsad18/go-supplier-otomax/resource"

	repos "github.com/aarsad18/go-supplier-otomax/repo"
	ucTransaksi "github.com/aarsad18/go-supplier-otomax/usecase"
)

type Services struct {
	UCTransaksi ucTransaksi.ITransaksiUsecase
}

func NewServices(db *resource.DBConn) (*Services, error) {

	repoSupplier := repos.NewSupplierRepo(db)
	repoTransaction := repos.NewTransactionRepo(db)
	repoOtomax := repos.NewOtomaxRepo()

	ucTransaksi := ucTransaksi.NewTransaksiUsecase(ucTransaksi.RepoInit{SupplierRepo: repoSupplier, TransactionRepo: repoTransaction, OtomaxRepo: repoOtomax})

	return &Services{UCTransaksi: ucTransaksi}, nil
}
