package eosws

import "encoding/json"

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

		Code      string `json:"code"`
		Scope     string `json:"scope"`
		TableName string `json:"table_name"`
	} `json:"data"`
}

type TableDelta struct {
	CommonIn

	Data struct {
		BlockNum  uint32          `json:"block_num"`
		Op        string          `json:"op,omitempty"`
		Code      string          `json:"code,omitempty"`
		Scope     string          `json:"scope,omitempty"`
		TableName string          `json:"name,omitempty"`
		Key       string          `json:"key,omitempty"`
		Payer     string          `json:"payer,omitempty"`
		JSON      json.RawMessage `json:"json"`
		Hex       string          `json:"hex"`
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
