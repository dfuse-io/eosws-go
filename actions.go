package eosws

import (
	"encoding/json"
)

func init() {
	RegisterIncomingMessage("action_trace", ActionTrace{})
	RegisterOutgoingMessage("get_actions", GetActions{})
}

type GetActions struct {
	CommonOut

	Data struct {
		Receiver                 string `json:"receiver,omitempty"`
		Account                  string `json:"account"`
		ActionName               string `json:"action_name,omitempty"`
		WithDBOperations         bool   `json:"with_db_ops"`
		WithRAMCosts             bool   `json:"with_ram_costs"`
		WithDeferredTransactions bool   `json:"with_deferred"`
		WithInlineTraces         bool   `json:"with_inline_traces"`
	} `json:"data"`
}

type ActionTrace struct {
	CommonIn
	Data struct {
		BlockNum             uint32          `json:"block_num"`
		BlockID              string          `json:"block_id"`
		TransactionID        string          `json:"trx_id"`
		ActionIndex          int             `json:"idx"`
		ActionDepth          int             `json:"depth"`
		Trace                json.RawMessage `json:"trace"`
		DBOperations         json.RawMessage `json:"dbops,omitempty"`
		RAMConsumed          json.RawMessage `json:"rams,omitempty"`
		DeferredTransactions json.RawMessage `json:"dtrxs,omitempty"`
	} `json:"data"`
}
