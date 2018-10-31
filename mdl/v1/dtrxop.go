package mdl

import (
	"encoding/json"
	"fmt"
	"strings"

	eos "github.com/eoscanada/eos-go"
)

type DTrxOp struct {
	Operation    string          `json:"op"`
	ActionIndex  int             `json:"action_idx"`
	Sender       string          `json:"sender"`
	SenderID     string          `json:"sender_id"`
	Payer        string          `json:"payer"`
	PublishedAt  string          `json:"published_at"`
	DelayUntil   string          `json:"delay_until"`
	ExpirationAt string          `json:"expiration_at"`
	TrxID        string          `json:"trx_id"`
	Trx          json.RawMessage `json:"trx,omitempty"`
}

type ExtDTrxOp struct {
	SourceTransactionID string             `json:"src_trx_id"`
	BlockNum            uint32             `json:"block_num"`
	BlockID             string             `json:"block_id"`
	BlockTime           eos.BlockTimestamp `json:"block_time"`

	DTrxOp
}

func (d *DTrxOp) IsCreateOperation() bool {
	dOp := strings.ToLower(d.Operation)
	if dOp == "push_create" || dOp == "modify_create" || dOp == "create" {
		return true
	}
	return false
}

func (d *DTrxOp) IsCancelOperation() bool {
	dOp := strings.ToLower(d.Operation)
	if dOp == "modify_cancel" || dOp == "cancel" {
		return true
	}
	return false
}

func (d *DTrxOp) SignedTransaction() (transaction *eos.SignedTransaction, err error) {
	err = json.Unmarshal(d.Trx, &transaction)
	if err != nil {
		return nil, fmt.Errorf("unmarshall signedTransaction: %s", err)
	}
	return transaction, nil
}
