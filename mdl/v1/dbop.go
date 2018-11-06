package mdl

type DBOp struct {
	Op          string `json:"op,omitempty"`
	ActionIndex int    `json:"action_idx"`
	Account     string `json:"account,omitempty"`
	Table       string `json:"table,omitempty"`
	Scope       string `json:"scope,omitempty"`
	Key         string `json:"key,omitempty"`
	Old         *DBRow `json:"old,omitempty"`
	New         *DBRow `json:"new,omitempty"`
}

type DBRow struct {
	Payer string      `json:"payer,omitempty"`
	Hex   string      `json:"hex,omitempty"`
	JSON  interface{} `json:"json,omitempty"`
	Error string      `json:"error,omitempty"`
}
