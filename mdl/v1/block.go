package mdl

import eos "github.com/eoscanada/eos-go"

// BlockSummary is the dfuse summary information for a given block
type BlockSummary struct {
	ID               string           `json:"id"`
	BlockNum         uint32           `json:"block_num"`
	Irreversible     bool             `json:"irreversible"`
	Header           *eos.BlockHeader `json:"header"`
	TransactionCount int              `json:"transaction_count"`
	SiblingBlocks    []*BlockSummary  `json:"sibling_blocks"`
}
