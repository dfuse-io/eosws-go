package v0

import "github.com/eoscanada/eos-go"

type DBOp struct {
	Operation   string `json:"op"`
	ActionIndex int    `json:"action_idx"`
	OldPayer    string `json:"opayer,omitempty"`
	NewPayer    string `json:"npayer,omitempty"`
	TablePath   string `json:"path"`
	OldData     string `json:"old,omitempty"`
	NewData     string `json:"new,omitempty"`

	chunks []string
}

type RAMOp struct {
	ActionIndex int    `json:"action_idx"`
	EventID     string `json:"event_id"`
	Family      string `json:"family"`
	Action      string `json:"action"`
	Operation string     `json:"op"`
	Payer     string     `json:"payer"`
	Delta     eos.Int64  `json:"delta"`
	Usage     eos.Uint64 `json:"usage"` // new usage
}

type DTrxOp struct {
	Operation     string           `json:"op"`
	ActionIndex   int              `json:"action_idx"`
	Sender        string           `json:"sender"`
	SenderID      string           `json:"sender_id"`
	Payer         string           `json:"payer"`
	PublishedAt   string           `json:"published_at"`
	DelayUntil    string           `json:"delay_until"`
	ExpirationAt  string           `json:"expiration_at"`
	TransactionID string           `json:"trx_id"`
	Transaction   *eos.Transaction `json:"trx,omitempty"`
}

type TableOp struct {
	Operation   string `json:"op"`
	ActionIndex int    `json:"action_idx"`
	Payer       string `json:"payer"`
	Path        string `json:"path"`
	chunks []string
}