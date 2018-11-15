package eosws

import (
	"encoding/json"

	v1 "github.com/dfuse-io/eosws-go/mdl/v1"
)

func init() {
	RegisterOutgoingMessage("get_table_rows", GetTableRows{})
	RegisterIncomingMessage("table_delta", TableDelta{})
	RegisterIncomingMessage("table_snapshot", TableSnapshot{})
}

type GetTableRows struct {
	CommonOut

	Data struct {
		JSON    bool `json:"json,omitempty"`
		Verbose bool `json:"verbose,omitempty"`

		Code  string `json:"code"`
		Scope string `json:"scope"`
		Table string `json:"table"`
	} `json:"data"`
}

type TableDelta struct {
	CommonIn

	Data struct {
		BlockNum uint32   `json:"block_num"`
		DBOp     *v1.DBOp `json:"dbop"`
		Step     string   `json:"step"`
	} `json:"data"`
}

type TableSnapshot struct {
	CommonIn

	Data struct {
		BlockNum uint32            `json:"block_num"`
		Rows     []json.RawMessage `json:"rows"`
	} `json:"data"`
}

type TableSnapshotRow struct {
	Key  string          `json:"key"`
	Data json.RawMessage `json:"data"`
}
