package repo

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/aarsad18/go-supplier-otomax/model"
)

type IOtomaxRepo interface {
	RequestTransaction(supplier model.Supplier, trx model.Transaction) (model.SupplierResult, error)
}

type OtomaxRepo struct {
}

func NewOtomaxRepo() *OtomaxRepo {
	return &OtomaxRepo{}
}

func (r *OtomaxRepo) RequestTransaction(supplier model.Supplier, trx model.Transaction) (result model.SupplierResult, err error) {
	// Execute the statement

	var (
		status TrxStatus
		reffSN string
		msg    string
	)

	requestURL := fmt.Sprintf("%s/trx?product=%s&dest=%s&refID=%s&memberID=%s&password=%s&pin=%s&qty=1", supplier.Url, trx.SupplierProductCode, trx.CustID, trx.TrxID, supplier.Username, supplier.Signature, supplier.PIN)

	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {

		// Read the response body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
		}

		stringBody := string(body)

		log.Printf("supplier response : %+v\n", stringBody)

		status = PROCESS
		reffSN = NA
		msg = ""

		rgxSuccess := regexp.MustCompile(`success`)
		rgxFailed := regexp.MustCompile(`gagal`)

		successMatch := rgxSuccess.FindStringSubmatch(stringBody)
		failedMatch := rgxFailed.FindStringSubmatch(stringBody)

		if len(successMatch) > 0 {
			status = SUCCESS

			rgxSn := regexp.MustCompile(`Saldo (?P<SN>.*) \- `)
			snMatch := rgxSn.FindStringSubmatch(stringBody)
			if len(snMatch) > 0 {
				sn := rgxSn.SubexpIndex("SN")
				reffSN = snMatch[sn]
			}

		} else if len(failedMatch) > 0 {
			status = FAILED
			msg = ""
		}

		result.Status = status
		result.SN = reffSN
		result.Msg = msg

		log.Printf("result : %+v\n", result)

		return result, err
	}

	return result, err
}
