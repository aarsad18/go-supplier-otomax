package model

const NA = "n/a"

type TrxStatus string

const (
	SUCCESS TrxStatus = "SUCCESS"
	FAILED  TrxStatus = "FAILED"
	PROCESS TrxStatus = "PROCESS"
)
