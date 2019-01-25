package eosws

import "github.com/shrimpliu/eosws-go/mdl/v1"

func init() {
	RegisterOutgoingMessage("get_transaction_lifecycle", GetTrxLifecycle{})
	RegisterIncomingMessage("transaction_lifecycle", TrxLifecycle{})
}

type GetTrxLifecycle struct {
	CommonOut
	Data struct{
		ID 	string	`json:"id"`
	}
}

type TrxLifecycle struct {
	CommonIn
	Data struct{
		Lifecycle *mdl.TransactionLifecycle `json:"lifecycle"`
	} `json:"data"`
}