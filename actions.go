package eosws

import (
	"encoding/json"
)

func init() {
	RegisterIncomingMessage("action_trace", ActionTrace{})
	RegisterOutgoingMessage("get_action_traces", GetActionTraces{})
}

type GetActionTraces struct {
	CommonOut

	Data struct {
		Receiver         string `json:"receiver,omitempty"`
		Account          string `json:"account"`
		ActionName       string `json:"action_name,omitempty"`
		WithDBOps        bool   `json:"with_dbops"`
		WithRAMOps       bool   `json:"with_ramops"`
		WithDTrxOps      bool   `json:"with_dtrxops"`
		WithInlineTraces bool   `json:"with_inline_traces"`
	} `json:"data"`
}

type ActionTrace struct {
	CommonIn
	Data struct {
		BlockNum      uint32          `json:"block_num"`
		BlockID       string          `json:"block_id"`
		TransactionID string          `json:"trx_id"`
		ActionIndex   int             `json:"idx"`
		ActionDepth   int             `json:"depth"`
		Trace         json.RawMessage `json:"trace"`
		DBOps         json.RawMessage `json:"dbops,omitempty"`
		RAMOps        json.RawMessage `json:"ramops,omitempty"`
		DTrxOps       json.RawMessage `json:"dtrxops,omitempty"`
	} `json:"data"`
}
