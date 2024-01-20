package repo

import (
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
		status model.TrxStatus
		reffSN string
		msg    string
	)

	// requestURL := fmt.Sprintf("%s/trx?product=%s&dest=%s&refID=%s&memberID=%s&password=%s&pin=%s&qty=1", supplier.Url, trx.SupplierProductCode, trx.CustID, trx.TrxID, supplier.Username, supplier.Signature, supplier.PIN)

	requestURL := "http://192.168.17.122:3000/transaction"
	// requestURL := "http://127.0.0.1:3000/transaction"

	log.Printf("REQUEST: %s\n", requestURL)

	res, err := http.Get(requestURL)
	if err != nil {
		log.Printf("REQUEST ERROR: %s\n", err)
		return model.SupplierResult{
			Status: model.PROCESS,
			Msg:    err.Error(),
		}, err
	}

	defer res.Body.Close()

	status = model.PROCESS
	reffSN = model.NA
	msg = ""

	if res.StatusCode == 200 {

		// Read the response body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Printf("ERROR READING BODY: %v\n", err)
			return model.SupplierResult{
				Status: model.PROCESS,
				Msg:    err.Error(),
			}, err
		}

		stringBody := string(body)

		log.Printf("RESPONSE: %+v\n", stringBody)

		rgxSuccess := regexp.MustCompile(`success`)
		rgxFailed := regexp.MustCompile(`gagal`)

		successMatch := rgxSuccess.FindStringSubmatch(stringBody)
		failedMatch := rgxFailed.FindStringSubmatch(stringBody)

		if len(successMatch) > 0 {
			status = model.SUCCESS

			rgxSn := regexp.MustCompile(`Saldo (?P<SN>.*) \- `)
			snMatch := rgxSn.FindStringSubmatch(stringBody)
			if len(snMatch) > 0 {
				sn := rgxSn.SubexpIndex("SN")
				reffSN = snMatch[sn]
			}

		} else if len(failedMatch) > 0 {
			status = model.FAILED
			msg = ""
		}

		result.Status = status
		result.SN = reffSN
		result.Msg = msg

	} else {
		log.Printf("HTTP ERROR [%v]: %s\n", res.StatusCode, err)
	}

	return result, err
}
